package grpc

import (
	"context"
	"time"

	grpcconfig "github.com/hodgesds/dlg/config/grpc"
	"github.com/hodgesds/dlg/executor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type grpcExecutor struct{}

// New returns a new gRPC executor.
func New() executor.GRPC {
	return &grpcExecutor{}
}

// Execute implements the gRPC executor interface.
func (e *grpcExecutor) Execute(ctx context.Context, config *grpcconfig.Config) error {
	opts := []grpc.DialOption{}

	if config.Insecure {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	if config.MaxConcurrent != nil {
		opts = append(opts, grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(*config.MaxConcurrent),
		))
	}

	conn, err := grpc.DialContext(ctx, config.Target, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Add metadata if provided
	if len(config.Metadata) > 0 {
		md := metadata.New(config.Metadata)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	// Set timeout if provided
	if config.Timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *config.Timeout)
		defer cancel()
	}

	// Execute the configured number of calls
	for i := 0; i < config.Count; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Generic health check or connection test
			// In a real implementation, this would invoke the specific method
			// For now, we just test the connection
			state := conn.GetState()
			if state.String() == "SHUTDOWN" {
				return nil
			}
			time.Sleep(10 * time.Millisecond)
		}
	}

	return nil
}
