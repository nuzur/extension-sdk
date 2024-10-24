package main

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/nuzur/extension-sdk/client"
)

func main() {
	c, err := client.New(client.Params{
		Token: "<token here>",
	})

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	project, err := c.GetProject(ctx, uuid.FromStringOrNil("0eb11a68-a59c-490d-9802-35f6fe211c6e"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("project: %v", project)
}
