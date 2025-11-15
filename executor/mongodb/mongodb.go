package mongodb

import (
	"context"
	"time"

	mongoconfig "github.com/hodgesds/dlg/config/mongodb"
	"github.com/hodgesds/dlg/executor"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoExecutor struct{}

// New returns a new MongoDB executor.
func New() executor.MongoDB {
	return &mongoExecutor{}
}

// Execute implements the MongoDB executor interface.
func (e *mongoExecutor) Execute(ctx context.Context, config *mongoconfig.Config) error {
	clientOpts := options.Client().ApplyURI(config.URI)

	if config.MaxPoolSize != nil {
		clientOpts.SetMaxPoolSize(*config.MaxPoolSize)
	}

	if config.MinPoolSize != nil {
		clientOpts.SetMinPoolSize(*config.MinPoolSize)
	}

	if config.ConnectTimeout != nil {
		clientOpts.SetConnectTimeout(*config.ConnectTimeout)
	}

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	collection := client.Database(config.Database).Collection(config.Collection)

	// Execute the configured number of operations
	for i := 0; i < config.Count; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := e.executeOperation(ctx, collection, config); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *mongoExecutor) executeOperation(ctx context.Context, collection *mongo.Collection, config *mongoconfig.Config) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	switch config.Operation {
	case mongoconfig.OpInsert:
		if config.Document != nil {
			_, err := collection.InsertOne(ctxWithTimeout, config.Document)
			return err
		}
	case mongoconfig.OpFind:
		cursor, err := collection.Find(ctxWithTimeout, config.Filter)
		if err != nil {
			return err
		}
		defer cursor.Close(ctxWithTimeout)
		// Consume cursor
		for cursor.Next(ctxWithTimeout) {
			var result interface{}
			if err := cursor.Decode(&result); err != nil {
				return err
			}
		}
		return cursor.Err()
	case mongoconfig.OpUpdate:
		if config.Filter != nil && config.Update != nil {
			_, err := collection.UpdateOne(ctxWithTimeout, config.Filter, config.Update)
			return err
		}
	case mongoconfig.OpDelete:
		if config.Filter != nil {
			_, err := collection.DeleteOne(ctxWithTimeout, config.Filter)
			return err
		}
	case mongoconfig.OpCount:
		_, err := collection.CountDocuments(ctxWithTimeout, config.Filter)
		return err
	}

	return nil
}
