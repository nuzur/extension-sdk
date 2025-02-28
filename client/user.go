package client

import (
	"context"

	"github.com/nuzur/extension-sdk/proto_deps/gen"
)

func (c *Client) IsValidUser(ctx context.Context) (bool, error) {
	user, err := c.productClient.GetTokenUser(ctx, &gen.GetTokenUserRequest{})
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	return true, nil
}
