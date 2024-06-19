package metrics

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func initExporter(collectorEndpoint string) (*otlptrace.Exporter, error) {
	traceExporter, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint(collectorEndpoint))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	return traceExporter, nil
}

func InitTracer(serviceName, collectorURL string) (*sdktrace.TracerProvider, error) {
	traceExporter, err := initExporter(collectorURL)
	if err != nil {
		return nil, err
	}

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
	)

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return tp, nil
}

func propagateHeaders(a *http.Request, b *http.Request) {
	headers := []string{
		"portal",
		"device",
		"user",
		"travel",
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"x-ot-span-context",
	}
	for _, header := range headers {
		b.Header.Add(header, a.Header.Get(header))
	}
}
