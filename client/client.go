package client

import (
	"crypto/x509"
	"errors"

	"github.com/gofrs/uuid"
	pb "github.com/nuzur/extension-sdk/proto_deps/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn          *grpc.ClientConn
	productClient pb.NuzurProductClient

	ExtensionUUID        string // this needs to be created before hand
	ExtensionVersionUUID string // this needs to be created before hand
	ExtensionName        string // returned by the metadata endpoint to identify the extension
	ExtensionAuthor      string // returned by the metadata endpoint to identify the extension
	NumberOfSteps        int32  // returned by the metadata endpoint
}

type Params struct {
	ExtensionUUID        string // this needs to be created before hand
	ExtensionVersionUUID string // this needs to be created before hand
	ExtensionName        string // returned by the metadata endpoint to identify the extension
	ExtensionAuthor      string // returned by the metadata endpoint to identify the extension
	NumberOfSteps        int32  // returned by the metadata endpoint

	API_ADDRESS *string // to enable testing
	DisableTLS  bool    // to enable testing
}

func New(params Params) (*Client, error) {

	if uuid.FromStringOrNil(params.ExtensionUUID) == uuid.Nil {
		return nil, errors.New("please provide a valid extension uuid")
	}
	if uuid.FromStringOrNil(params.ExtensionVersionUUID) == uuid.Nil {
		return nil, errors.New("please provide a valid extension version uuid")
	}

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

	// TODO possibly call api to get the extension details to validate

	return &Client{
		conn:          conn,
		productClient: productClient,

		ExtensionUUID:        params.ExtensionUUID,
		ExtensionVersionUUID: params.ExtensionVersionUUID,
		ExtensionName:        params.ExtensionName,
		ExtensionAuthor:      params.ExtensionAuthor,
		NumberOfSteps:        params.NumberOfSteps,
	}, nil
}
