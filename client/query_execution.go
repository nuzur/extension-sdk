package client

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/proto_deps/gen"
)

func (c *Client) GetQueryExecution(ctx context.Context, queryExecutionUUID uuid.UUID) (*gen.GetQueryExecutionResponse, error) {
	return c.connectionManagerClient.GetQueryExecution(ctx, &gen.GetQueryExecutionRequest{
		ExecutionUuid: queryExecutionUUID.String(),
	})
}

func (c *Client) ExecuteQuery(ctx context.Context, req *gen.ExecuteQueryRequest) (*gen.ExecuteQueryResponse, error) {
	return c.connectionManagerClient.ExecuteQuery(ctx, req)
}

func (c *Client) CancelQueryExecution(ctx context.Context, queryExecutionUUID uuid.UUID) (*gen.CancelQueryExecutionResponse, error) {
	return c.connectionManagerClient.CancelQueryExecution(ctx, &gen.CancelQueryExecutionRequest{
		ExecutionUuid: queryExecutionUUID.String(),
	})
}
