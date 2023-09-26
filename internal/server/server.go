package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/ilhame90/gymshark-task/internal/config"
	"github.com/ilhame90/gymshark-task/internal/custom_middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// New creates new echo instance
func New(lc fx.Lifecycle, shutdowner fx.Shutdowner, cfg *config.Config, log *zap.Logger) (*echo.Echo, error) {
	engine := echo.New()

	// Setup middlewares
	engine.Use(custom_middleware.ZapLoggerMiddleware(log))
	engine.Use(middleware.CORSWithConfig(cfg.HttpServerConfig.CORSConfig))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				addr := fmt.Sprintf(":%s", cfg.HttpServerConfig.Port)
				listener, err := net.Listen("tcp", addr)
				if err != nil {
					engine.Logger.Fatal("shutting down the server is started due to listener", err.Error())
					shutdowner.Shutdown()
					return
				}
				engine.Listener = listener

				if err := engine.Start(addr); err != nil && err != http.ErrServerClosed {
					engine.Logger.Fatal("shutting down the server is started due to server")
					shutdowner.Shutdown()
					return
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// Graceful shutdown
			engine.Logger.Info("shutting down the server gracefully")
			return engine.Shutdown(ctx)
		},
	})

	return engine, nil
}
