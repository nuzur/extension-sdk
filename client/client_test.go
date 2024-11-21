package client

import "github.com/nuzur/extension-sdk/dummyserver"

func BuildDummyClientServer() (*Client, *dummyserver.Server, error) {
	server, err := dummyserver.New()
	if err != nil {
		return nil, nil, err
	}

	client, err := New(Params{
		API_ADDRESS: server.Address(),
		DisableTLS:  true,
	})

	return client, server, err
}
