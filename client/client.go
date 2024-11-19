package client

import (
	"crypto/x509"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/nuzur/extension-sdk/config"
	pb "github.com/nuzur/extension-sdk/proto_deps/gen"
	"go.uber.org/fx"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn          *grpc.ClientConn
	productClient pb.NuzurProductClient
	metadata      *config.ExtensionConfigMetadata
	i18nbundle    *i18n.Bundle
}

type Params struct {
	fx.In

	ConfigPath  *string `optional:"true"`
	API_ADDRESS *string `optional:"true"`
	DisableTLS  bool    `optional:"true"`
}

func New(params Params) (*Client, error) {
	// load config
	configPath := filepath.Join(CurrentPath(), "config")
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

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	translationPath := filepath.Join(CurrentPath(), "translations")
	bundle.LoadMessageFile(filepath.Join(translationPath, "en.toml"))
	bundle.LoadMessageFile(filepath.Join(translationPath, "es.toml"))

	// build grpc client
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
		i18nbundle:    bundle,
	}, nil
}

func CurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
