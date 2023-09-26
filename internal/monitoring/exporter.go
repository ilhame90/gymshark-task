package monitoring

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
)

func NewOtelFileExporter(lc fx.Lifecycle) (trace.SpanExporter, error) {
	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	f, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal(err)
	}

	lc.Append(fx.Hook{
		OnStop: func(c context.Context) error {
			return f.Close()
		},
	})

	return stdouttrace.New(
		stdouttrace.WithWriter(f),
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)
}
