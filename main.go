package main

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"log"

	"github.com/zenbusiness/go-toolkit/app"
	"github.com/zenbusiness/go-toolkit/env"
	internalgrpc "github.com/zenbusiness/go-toolkit/grpc"
	grpcinterceptor "github.com/zenbusiness/go-toolkit/grpc/interceptor"
	"github.com/zenbusiness/go-toolkit/launchdarkly"
	ldinterceptor "github.com/zenbusiness/go-toolkit/launchdarkly/interceptor"
	"github.com/zenbusiness/go-toolkit/logging"
	logctx "github.com/zenbusiness/go-toolkit/logging/context"
	loginterceptor "github.com/zenbusiness/go-toolkit/logging/interceptor"
	"github.com/zenbusiness/go-toolkit/redis"
	"github.com/zenbusiness/go-toolkit/telemetry"
	"github.com/zenbusiness/go-toolkit/telemetry/metrics"
	"github.com/zenbusiness/go-toolkit/telemetry/tracing"
	"github.com/zenbusiness/golang-service-template/internal/config"
	"github.com/zenbusiness/golang-service-template/internal/service"
	hw "github.com/zenbusiness/proto-go/gen/zenbusiness/misc/v1alpha1"
)

// main is just an example of a way to pull in all the libraries
func main() {
	conf := config.Config{}
	err := env.Load(context.Background(), &conf)
	if err != nil {
		log.Fatal("failed to parse configuration", zap.Error(err))
	}

	myApp, err := initialize(&conf)
	if err != nil {
		log.Fatal("failed to initialize app", zap.Error(err))
	}

	if err := myApp.Run(context.Background()); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal("failure in app", zap.Error(err))
	}
}

func initialize(conf *config.Config) (*app.App, error) {
	var deps []app.Dependency
	logger, err := logging.CreateLogger(conf.Logging, conf.Environment)
	if err != nil {
		return nil, err
	}
	deps = append(deps, logger)

	metricRegistry := metrics.NewMetricsRegistry(conf.Metrics)
	telemetryServ, err := telemetry.CreateTelemetryServer(conf.Telemetry, telemetry.WithRegistry(metricRegistry))
	if err != nil {
		return nil, err
	}
	deps = append(deps, telemetryServ)

	tracer, err := tracing.NewTracerProvider(conf.Tracing)
	if err != nil {
		return nil, err
	}
	deps = append(deps, tracer)

	ldClient, err := launchdarkly.NewLaunchDarkly(conf.LD, metricRegistry)
	if err != nil {
		return nil, err
	}
	deps = append(deps, ldClient)

	redisClient, err := redis.NewRedisClient(conf.Redis, metricRegistry)
	if err != nil {
		return nil, err
	}
	deps = append(deps, redisClient)

	myServer, err := internalgrpc.CreateGRPCServer(conf.Grpc,
		internalgrpc.WithHealthChecks(telemetryServ, ldClient, redisClient),
		internalgrpc.WithUnaryInterceptors(
			ldinterceptor.NewLDContextInterceptor(metricRegistry),
			grpcinterceptor.NewMetricsInterceptor(metricRegistry),
			loginterceptor.RequestResponseInterceptor(logctx.GetCtxLogger(context.Background()))),
	)
	if err != nil {
		return nil, err
	}

	s := service.Server{}
	hw.RegisterHelloWorldServiceServer(myServer.GrpcServer(), &s)
	deps = append(deps, myServer)

	// Create the app after all dependencies have been gathered
	myApp, err := app.CreateApp(conf.App, deps...)
	if err != nil {
		return nil, err
	}

	return myApp, nil
}
