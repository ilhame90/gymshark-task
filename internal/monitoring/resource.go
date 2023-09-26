package monitoring

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// newResource returns a resource describing this application.
func newOtelResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("orders"),
			semconv.ServiceVersionKey.String("v1"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}
