package neo4j

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	neo4jconfig "github.com/hodgesds/dlg/config/neo4j"
)

type neo4jExecutor struct{}

func New() *neo4jExecutor {
	return &neo4jExecutor{}
}

func (e *neo4jExecutor) Execute(ctx context.Context, config *neo4jconfig.Config) error {
	if err := config.Validate(); err != nil {
		return err
	}

	driver, err := neo4j.NewDriverWithContext(config.URI, neo4j.BasicAuth(config.Username, config.Password, ""))
	if err != nil {
		return fmt.Errorf("failed to create driver: %w", err)
	}
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: config.Database})
	defer session.Close(ctx)

	_, err = session.Run(ctx, config.Query, config.Parameters)
	return err
}
