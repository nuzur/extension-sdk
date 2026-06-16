package extensionsdk

import (
	"time"

	"github.com/nuzur/extension-sdk/client"
	"github.com/nuzur/extension-sdk/server"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("extensionsdk",
	fx.StopTimeout(45 * time.Second),
	fx.Provide(
		zap.NewProduction,
		client.New,
	),
	fx.Invoke(server.New),
)
