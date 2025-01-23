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
	conn          *grpc.ClientConn
	productClient pb.NuzurProductClient
	metadata      *config.ExtensionConfigMetadata
	i18nbundle    *i18n.Bundle
}

type Params struct {
	fx.In

	ConfigPath          *string `optional:"true"`
	PRODUCT_API_ADDRESS *string `optional:"true"`
	DisableTLS          bool    `optional:"true"`
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

	api_address := PRODUCT_API_PROD_ADDRESS
	if params.PRODUCT_API_ADDRESS != nil {
		api_address = *params.PRODUCT_API_ADDRESS
	}
	conn, err := grpc.NewClient(api_address, opts...)
	if err != nil {
		return nil, err
	}

	productClient := pb.NewNuzurProductClient(conn)

	// TODO possibly call api to get the extension details to validate
	// might not be possible because we don't have auth token at this point

	return &Client{
		conn:          conn,
		productClient: productClient,
		metadata:      &metadata,
		i18nbundle:    bundle,
	}, nil
}
