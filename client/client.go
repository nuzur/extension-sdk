package client

import (
	"crypto/x509"
	"path/filepath"
	"runtime"

	"github.com/nuzur/extension-sdk/config"
	pb "github.com/nuzur/extension-sdk/proto_deps/gen"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn          *grpc.ClientConn
	productClient pb.NuzurProductClient
	metadata      *config.ExtensionConfigMetadata
}

type Params struct {
	fx.In

	ConfigPath  *string `optional:"true"`
	API_ADDRESS *string `optional:"true"`
	DisableTLS  bool    `optional:"true"`
}

func New(params Params) (*Client, error) {
	configPath := filepath.Join(RootPath(), "config")
	if params.ConfigPath != nil {
		configPath = *params.ConfigPath
	}
	configProvider, err := config.New(configPath)
	if err != nil {
		return nil, err
	}

	metadata := config.ExtensionConfigMetadata{}
	err = configProvider.Get(METADATA).Populate(&metadata)
	if err != nil {
		return nil, err
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
		metadata:      &metadata,
	}, nil
}

func RootPath() string {
	_, callerFile, _, _ := runtime.Caller(1)
	generatorDir := filepath.Dir(callerFile)
	absoluteGeneratorDir, err := filepath.Abs(generatorDir)
	if err != nil {
		panic("could not resolve path")
	}
	return filepath.Dir(absoluteGeneratorDir)
}
