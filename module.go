package extensionsdk

import (
	"github.com/nuzur/extension-sdk/client"
	"github.com/nuzur/extension-sdk/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Provide(
	zap.NewProduction,
	client.New,
	fx.Invoke(server.New),
)
