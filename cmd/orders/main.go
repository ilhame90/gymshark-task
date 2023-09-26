package main

import (
	"github.com/ilhame90/gymshark-task/internal/config"
	"github.com/ilhame90/gymshark-task/internal/logger"
	"github.com/ilhame90/gymshark-task/internal/monitoring"
	"github.com/ilhame90/gymshark-task/internal/orders"
	"github.com/ilhame90/gymshark-task/internal/orders/delivery/http"
	"github.com/ilhame90/gymshark-task/internal/orders/usecase"
	"github.com/ilhame90/gymshark-task/internal/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.New,
			logger.New,
			server.New,
			fx.Annotate(
				usecase.NewOrdersUsecase,
				fx.As(new(orders.Usecase)),
			),
			http.NewOrdersHandler,
			monitoring.NewOtelFileExporter,
			monitoring.NewOtelTracerProvider,
		),
		fx.Invoke(monitoring.ResgisterTracerProvider, server.RegisterRoutes),
	)

	app.Run()
}
