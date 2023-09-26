package monitoring

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
)

func NewOtelTracerProvider(lc fx.Lifecycle, exp trace.SpanExporter) *trace.TracerProvider {
	l := log.New(os.Stdout, "", 0)

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newOtelResource()),
	)

	lc.Append(fx.Hook{
		OnStop: func(c context.Context) error {
			if err := tp.Shutdown(context.Background()); err != nil {
				l.Fatal(err)
				return err
			}
			return nil
		},
	})

	return tp
}
