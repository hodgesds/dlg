package graphql

import (
	"context"

	graphqlconfig "github.com/hodgesds/dlg/config/graphql"
	"github.com/hodgesds/dlg/executor"
	"github.com/machinebox/graphql"
)

type graphqlExecutor struct{}

// New returns a new GraphQL executor.
func New() executor.GraphQL {
	return &graphqlExecutor{}
}

// Execute implements the GraphQL executor interface.
func (e *graphqlExecutor) Execute(ctx context.Context, config *graphqlconfig.Config) error {
	client := graphql.NewClient(config.Endpoint)

	// Execute the configured number of queries
	for i := 0; i < config.Count; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			req := graphql.NewRequest(config.Query)

			// Add variables if provided
			if config.Variables != nil {
				for key, value := range config.Variables {
					req.Var(key, value)
				}
			}

			// Add headers if provided
			if config.Headers != nil {
				for key, value := range config.Headers {
					req.Header.Set(key, value)
				}
			}

			var response interface{}

			// Apply timeout if configured
			execCtx := ctx
			if config.Timeout != nil {
				var cancel context.CancelFunc
				execCtx, cancel = context.WithTimeout(ctx, *config.Timeout)
				defer cancel()
			}

			if err := client.Run(execCtx, req, &response); err != nil {
				return err
			}
		}
	}

	return nil
}
