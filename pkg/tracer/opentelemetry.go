package tracer

import (
	"context"
	"github.com/harisaginting/tech-test-adsi/pkg/log"
	"github.com/harisaginting/tech-test-adsi/pkg/utils/helper"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/exporters/jaeger"

)

var span oteltrace.Span
var tracer oteltrace.Tracer

func InitTracer() {
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// jaeger
	exp, err := jaeger.New(
		jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost("localhost"),
			jaeger.WithAgentPort("6831"),
		),
	)
	// stdout
	// exp, err := stdout.New(stdout.WithPrettyPrint())
	
	if err != nil {
		return
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("ginting-srv"),
		)),
	)
	if err != nil {
		return
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	tracer = tp.Tracer("ginting-app")
}

func Span(ctx context.Context, name string) oteltrace.Span {
	log.Trace(ctx,name)
	_, span = tracer.Start(ctx,name)
	return span
}

func SetAttributeString(span oteltrace.Span,key string, value interface{}){
	val := helper.ForceString(value)
	span.SetAttributes(attribute.String(key, val))
}

func SetAttributeInt(span oteltrace.Span,key string, value interface{}){
	val := helper.ForceInt(value)
	span.SetAttributes(attribute.Int(key, val))
}

func addEvent(span oteltrace.Span,event string){
	span.AddEvent(event)
}