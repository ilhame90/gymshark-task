package monitoring

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func ResgisterTracerProvider(tp *trace.TracerProvider) {
	otel.SetTracerProvider(tp)
}
