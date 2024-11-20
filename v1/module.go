package extensionsdk

import (
	"github.com/nuzur/extension-sdk/v1/client"
	"github.com/nuzur/extension-sdk/v1/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("extensionsdk",
	fx.Provide(
		zap.NewProduction,
		client.New,
	),
	fx.Invoke(server.New),
)
