package client

import (
	"crypto/x509"

	pb "github.com/nuzur/extension-sdk/proto_deps/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	token         string
	conn          *grpc.ClientConn
	productClient pb.NuzurProductClient
}

type Params struct {
	Token       string
	API_ADDRESS *string
	DisableTLS  bool
}

func New(params Params) (*Client, error) {
	var opts []grpc.DialOption

	if params.DisableTLS {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		pool, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		creds := credentials.NewClientTLSFromCert(pool, "")
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	api_address := API_PROD_ADDRESS
	if params.API_ADDRESS != nil {
		api_address = *params.API_ADDRESS
	}
	conn, err := grpc.NewClient(api_address, opts...)
	if err != nil {
		return nil, err
	}

	productClient := pb.NewNuzurProductClient(conn)

	return &Client{
		token:         params.Token,
		conn:          conn,
		productClient: productClient,
	}, nil
}
