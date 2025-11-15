package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	elasticsearchconfig "github.com/hodgesds/dlg/config/elasticsearch"
	"github.com/hodgesds/dlg/executor"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type elasticsearchExecutor struct{}

// New returns a new Elasticsearch executor.
func New() executor.Elasticsearch {
	return &elasticsearchExecutor{}
}

// Execute implements the Elasticsearch executor interface.
func (e *elasticsearchExecutor) Execute(ctx context.Context, config *elasticsearchconfig.Config) error {
	cfg := elasticsearch.Config{
		Addresses: config.Addresses,
	}

	if config.Username != "" && config.Password != "" {
		cfg.Username = config.Username
		cfg.Password = config.Password
	}

	if config.CloudID != "" {
		cfg.CloudID = config.CloudID
	}

	if config.APIKey != "" {
		cfg.APIKey = config.APIKey
	}

	if config.MaxRetries != nil {
		cfg.MaxRetries = *config.MaxRetries
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}

	// Execute the configured number of operations
	for i := 0; i < config.Count; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := e.executeOperation(ctx, client, config); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *elasticsearchExecutor) executeOperation(ctx context.Context, client *elasticsearch.Client, config *elasticsearchconfig.Config) error {
	switch config.Operation {
	case elasticsearchconfig.OpIndex:
		if config.Document != nil {
			doc, err := json.Marshal(config.Document)
			if err != nil {
				return err
			}

			req := esapi.IndexRequest{
				Index:      config.Index,
				DocumentID: config.DocumentID,
				Body:       bytes.NewReader(doc),
				Refresh:    "false",
			}

			res, err := req.Do(ctx, client)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			if res.IsError() {
				return fmt.Errorf("error indexing document: %s", res.Status())
			}
		}

	case elasticsearchconfig.OpGet:
		req := esapi.GetRequest{
			Index:      config.Index,
			DocumentID: config.DocumentID,
		}

		res, err := req.Do(ctx, client)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.IsError() {
			return fmt.Errorf("error getting document: %s", res.Status())
		}

	case elasticsearchconfig.OpSearch:
		if config.Query != nil {
			query, err := json.Marshal(map[string]interface{}{
				"query": config.Query,
			})
			if err != nil {
				return err
			}

			req := esapi.SearchRequest{
				Index: []string{config.Index},
				Body:  bytes.NewReader(query),
			}

			res, err := req.Do(ctx, client)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			if res.IsError() {
				return fmt.Errorf("error searching: %s", res.Status())
			}
		}

	case elasticsearchconfig.OpUpdate:
		if config.Document != nil {
			doc, err := json.Marshal(map[string]interface{}{
				"doc": config.Document,
			})
			if err != nil {
				return err
			}

			req := esapi.UpdateRequest{
				Index:      config.Index,
				DocumentID: config.DocumentID,
				Body:       bytes.NewReader(doc),
			}

			res, err := req.Do(ctx, client)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			if res.IsError() {
				return fmt.Errorf("error updating document: %s", res.Status())
			}
		}

	case elasticsearchconfig.OpDelete:
		req := esapi.DeleteRequest{
			Index:      config.Index,
			DocumentID: config.DocumentID,
		}

		res, err := req.Do(ctx, client)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.IsError() && !strings.Contains(res.Status(), "404") {
			return fmt.Errorf("error deleting document: %s", res.Status())
		}
	}

	return nil
}
