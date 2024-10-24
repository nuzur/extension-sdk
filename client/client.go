package client

import (
	"crypto/x509"

	pb "github.com/nuzur/extension-sdk/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	token         string
	conn          *grpc.ClientConn
	productClient pb.NuzurProductClient
}

type Params struct {
	Token string
}

func New(params Params) (*Client, error) {
	var opts []grpc.DialOption

	pool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	creds := credentials.NewClientTLSFromCert(pool, "")
	opts = append(opts, grpc.WithTransportCredentials(creds))

	// TODO move address to config
	conn, err := grpc.NewClient("api.nuzur.com:443", opts...)
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
