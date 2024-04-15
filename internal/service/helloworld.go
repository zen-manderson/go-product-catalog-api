package service

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ldctx "github.com/zenbusiness/go-toolkit/logging/context"
	hw "github.com/zenbusiness/proto-go/gen/zenbusiness/misc/v1alpha1"
)

// Ensure interface compliance.
var _ hw.HelloWorldServiceServer = (*Server)(nil)

// Server is a type implementation of the helloworld RPC from
// the zenbusiness.misc.v1alpha1 proto file
type Server struct {
	hw.UnimplementedHelloWorldServiceServer
}

// SayHello is a receiver implementation off of the helloworld RPC.
//
// https://github.com/zenbusiness/proto-registry/blob/f1498de75120f0c8fca5bf95e5a152d7809caa89/proto/zenbusiness/misc/v1alpha1/helloworld.proto#L11
//
// The name is very specific to the definition of `SayHelloRequest` - just remove Request.
// When an RPC is received - the message being sent must be called `SayHello`. The gRPC service will then push the message into
// the implemented receiver. This is where all business logic will go for any RPC.
func (s *Server) SayHello(ctx context.Context, in *hw.SayHelloRequest) (*hw.SayHelloResponse, error) {
	logger := ldctx.GetCtxLogger(ctx)
	// Logging request payloads is expensive and noisy. It is however useful in the dev and local environment.
	logger.Debug("SayHello: start", zap.Any("req", in))

	// Always validate the request before doing anything else
	if in.GetName() == "" {
		statusCode := status.Error(codes.InvalidArgument, "name is required")
		logger.Error("invalid argument", zap.Error(statusCode))
		return nil, statusCode
	}

	// Add a field to the logger
	logger = logger.With(zap.String("name", in.GetName()))
	// inject the new logger so subsequent logs have the field "name"
	ctx = ldctx.ToContext(ctx, logger)

	res := &hw.SayHelloResponse{
		Message: "Hello, " + in.Name,
	}

	// Logging response payloads is expensive and noisy. It is however useful in the dev and local environment.
	logger.Debug("SayHello: end", zap.Any("res", res))

	return res, nil
}
