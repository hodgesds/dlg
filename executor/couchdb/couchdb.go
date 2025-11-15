package couchdb

import (
	"context"
	"fmt"
	"net/http"

	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
	couchconfig "github.com/hodgesds/dlg/config/couchdb"
)

type couchExecutor struct{}

func New() *couchExecutor {
	return &couchExecutor{}
}

func (e *couchExecutor) Execute(ctx context.Context, config *couchconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	httpClient := &http.Client{Timeout: config.Timeout}
	client, err := kivik.New("couch", config.URL, kivik.WithHTTPClient(httpClient))
	if err != nil {
		return err
	}

	if config.Username != "" && config.Password != "" {
		err = client.Authenticate(ctx, kivik.BasicAuth(config.Username, config.Password))
		if err != nil {
			return err
		}
	}

	db := client.DB(config.Database)

	switch config.Operation {
	case "get":
		if config.DocumentID == "" {
			return fmt.Errorf("document_id required for get")
		}
		var doc interface{}
		return db.Get(ctx, config.DocumentID).ScanDoc(&doc)

	case "put":
		if config.DocumentID == "" || config.Document == "" {
			return fmt.Errorf("document_id and document required for put")
		}
		_, err := db.Put(ctx, config.DocumentID, config.Document)
		return err

	case "delete":
		if config.DocumentID == "" {
			return fmt.Errorf("document_id required for delete")
		}
		var doc map[string]interface{}
		err := db.Get(ctx, config.DocumentID).ScanDoc(&doc)
		if err != nil {
			return err
		}
		rev, ok := doc["_rev"].(string)
		if !ok {
			return fmt.Errorf("failed to get revision")
		}
		_, err = db.Delete(ctx, config.DocumentID, rev)
		return err

	case "query":
		if config.Query == "" {
			return fmt.Errorf("query required for query operation")
		}
		rows := db.Find(ctx, config.Query)
		defer rows.Close()
		if rows.Next() {
			var doc interface{}
			return rows.ScanDoc(&doc)
		}
		return rows.Err()

	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}
}
