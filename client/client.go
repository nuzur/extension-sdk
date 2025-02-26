package client

import (
	"crypto/x509"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/nuzur/extension-sdk/config"
	pb "github.com/nuzur/extension-sdk/proto_deps/gen"
	"github.com/nuzur/filetools"
	"go.uber.org/fx"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	productConn             *grpc.ClientConn
	connectionManagerConn   *grpc.ClientConn
	productClient           pb.NuzurProductClient
	connectionManagerClient pb.NuzurConnectionManagerClient
	metadata                *config.ExtensionConfigMetadata
	i18nbundle              *i18n.Bundle
}

type Params struct {
	fx.In

	ConfigPath                     *string `optional:"true"`
	PRODUCT_API_ADDRESS            *string `optional:"true"`
	CONNECTION_MANAGER_API_ADDRESS *string `optional:"true"`
	DisableTLS                     bool    `optional:"true"`
}

func New(params Params) (*Client, error) {
	// load config
	configPath := filepath.Join(filetools.CurrentLocalPath(), "config")
	if params.ConfigPath != nil {
		configPath = *params.ConfigPath
	}
	configProvider, err := config.New(configPath)
	if err != nil {
		return nil, err
	}

	// read metadata
	metadata := config.ExtensionConfigMetadata{}
	err = configProvider.Get(METADATA).Populate(&metadata)
	if err != nil {
		return nil, err
	}

	// translations
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	translationPath := filepath.Join(filetools.CurrentLocalPath(), "translations")
	bundle.LoadMessageFile(filepath.Join(translationPath, "en.toml"))
	bundle.LoadMessageFile(filepath.Join(translationPath, "es.toml"))

	// build grpc clients
	productClient, productConn, err := buildProductClient(params)
	if err != nil {
		return nil, err
	}

	connectionManagerClient, connectionManagerConn, err := buildConnectioManagerClient(params)
	if err != nil {
		return nil, err
	}

	// TODO possibly call api to get the extension details to validate
	// might not be possible because we don't have auth token at this point

	return &Client{
		productConn:             productConn,
		productClient:           productClient,
		connectionManagerConn:   connectionManagerConn,
		connectionManagerClient: connectionManagerClient,
		metadata:                &metadata,
		i18nbundle:              bundle,
	}, nil
}

func buildProductClient(params Params) (pb.NuzurProductClient, *grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if params.DisableTLS {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		pool, err := x509.SystemCertPool()
		if err != nil {
			return nil, nil, err
		}
		creds := credentials.NewClientTLSFromCert(pool, "")
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	api_address := PRODUCT_API_PROD_ADDRESS
	if params.PRODUCT_API_ADDRESS != nil {
		api_address = *params.PRODUCT_API_ADDRESS
	}

	conn, err := grpc.NewClient(api_address, opts...)
	if err != nil {
		return nil, nil, err
	}
	productClient := pb.NewNuzurProductClient(conn)
	return productClient, conn, nil
}

func buildConnectioManagerClient(params Params) (pb.NuzurConnectionManagerClient, *grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if params.DisableTLS {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		pool, err := x509.SystemCertPool()
		if err != nil {
			return nil, nil, err
		}
		creds := credentials.NewClientTLSFromCert(pool, "")
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	api_address := CONNECTION_MANAGER_API_PROD_ADDRESS
	if params.CONNECTION_MANAGER_API_ADDRESS != nil {
		api_address = *params.CONNECTION_MANAGER_API_ADDRESS
	}

	conn, err := grpc.NewClient(api_address, opts...)
	if err != nil {
		return nil, nil, err
	}
	productClient := pb.NewNuzurConnectionManagerClient(conn)
	return productClient, conn, nil
}
