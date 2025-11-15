package arangodb

import (
	"context"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	arangoconfig "github.com/hodgesds/dlg/config/arangodb"
)

type arangoExecutor struct{}

func New() *arangoExecutor {
	return &arangoExecutor{}
}

func (e *arangoExecutor) Execute(ctx context.Context, config *arangoconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	conn, err := http.NewConnection(http.ConnectionConfig{Endpoints: config.Endpoints})
	if err != nil {
		return err
	}

	var client driver.Client
	if config.Username != "" && config.Password != "" {
		client, err = driver.NewClient(driver.ClientConfig{
			Connection:     conn,
			Authentication: driver.BasicAuthentication(config.Username, config.Password),
		})
	} else {
		client, err = driver.NewClient(driver.ClientConfig{Connection: conn})
	}
	if err != nil {
		return err
	}

	db, err := client.Database(ctx, config.Database)
	if err != nil {
		return err
	}

	switch config.Operation {
	case "query":
		if config.Query == "" {
			return fmt.Errorf("query required")
		}
		cursor, err := db.Query(ctx, config.Query, config.BindVars)
		if err != nil {
			return err
		}
		defer cursor.Close()
		var doc interface{}
		_, err = cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			return nil
		}
		return err

	case "insert":
		if config.Collection == "" || config.Document == nil {
			return fmt.Errorf("collection and document required")
		}
		col, err := db.Collection(ctx, config.Collection)
		if err != nil {
			return err
		}
		_, err = col.CreateDocument(ctx, config.Document)
		return err

	case "update":
		if config.Collection == "" || config.Key == "" || config.Document == nil {
			return fmt.Errorf("collection, key, and document required")
		}
		col, err := db.Collection(ctx, config.Collection)
		if err != nil {
			return err
		}
		_, err = col.UpdateDocument(ctx, config.Key, config.Document)
		return err

	case "delete":
		if config.Collection == "" || config.Key == "" {
			return fmt.Errorf("collection and key required")
		}
		col, err := db.Collection(ctx, config.Collection)
		if err != nil {
			return err
		}
		_, err = col.RemoveDocument(ctx, config.Key)
		return err

	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}
}
