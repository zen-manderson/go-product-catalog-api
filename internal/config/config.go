package config

import (
	"github.com/zenbusiness/go-toolkit/app"
	internalgrpc "github.com/zenbusiness/go-toolkit/grpc"
	"github.com/zenbusiness/go-toolkit/launchdarkly"
	"github.com/zenbusiness/go-toolkit/logging"
	"github.com/zenbusiness/go-toolkit/redis"
	"github.com/zenbusiness/go-toolkit/telemetry"
	"github.com/zenbusiness/go-toolkit/telemetry/metrics"
	"github.com/zenbusiness/go-toolkit/telemetry/tracing"
)

// Config is the struct that will hold all of your app's configuration. It is best to keep the config fields segmented by
// package.
type Config struct {
	// Environment is required. If this is missing your app will fail to start up
	Environment string `env:"ENVIRONMENT,required"`
	// Telemetry controls how metrics are exported and served
	Telemetry telemetry.Config `env:",prefix=TELEMETRY_"`
	// Metrics controls what metrics are collected
	Metrics metrics.Config `env:",prefix=METRICS_"`
	// Logging contains all configurations specific to logging. All environment variables will be prefixed with "LOG_"
	Logging logging.Config `env:",prefix=LOG_"`
	// Grpc contains all configurations specific to grpc. All environment variables will be prefixed with "GRPC_"
	Grpc internalgrpc.Config `env:",prefix=GRPC_"`
	// LD contains all configurations specific to launch darkly. All environment variables will be prefixed with "LD_"
	LD launchdarkly.Config `env:",prefix=LD_"`
	// Redis contains all configurations specific to redis. All environment variables will be prefixed with "REDIS_"
	Redis redis.Config `env:",prefix=REDIS_"`
	// Tracing controls how traces are exported
	Tracing tracing.Config `env:",prefix=TRACING_"`
	// App contains environment variables for controlling how app package behaves
	App app.Config `env:",prefix=APP_"`
}
