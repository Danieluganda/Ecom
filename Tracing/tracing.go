package main

/* Distributed Tracing: Trace id propagated across the call chain of a transaction to
understand how a single transaction/system behaves. */

import (
	"log"
	"net/http"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

func main() {
	// Initialize Jaeger tracer
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jaeger tracer: %v", err)
	}
	cfg.ServiceName = "my-service"
	cfg.Sampler.Type = jaeger.SamplerTypeConst
	cfg.Sampler.Param = 1
	cfg.Reporter.LogSpans = true
	cfg.Reporter.BufferFlushInterval = time.Second
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatalf("Failed to create Jaeger tracer: %v", err)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/", handleRequest)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := opentracing.GlobalTracer().StartSpan("handleRequest", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	// Simulating some work
	time.Sleep(100 * time.Millisecond)
	span.LogFields(log.String("event", "work done"))

	childSpan := opentracing.GlobalTracer().StartSpan("additionalWork", opentracing.ChildOf(span.Context()))
	defer childSpan.Finish()

	// Simulating additional work
	time.Sleep(50 * time.Millisecond)
	childSpan.LogFields(log.String("event", "additional work done"))

	w.Write([]byte("Hello, World!"))
}
