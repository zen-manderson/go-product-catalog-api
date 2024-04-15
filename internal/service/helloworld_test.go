package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"

	"github.com/stretchr/testify/assert"
	hw "github.com/zenbusiness/proto-go/gen/zenbusiness/misc/v1alpha1"
	"go.uber.org/zap"
)

// TestSayHello is confirming our HelloWorld-like request works as expected with metadata
func TestSayHello(t *testing.T) {
	log := zap.L()
	ctx := context.Background()
	updatedCtx := context.WithValue(ctx, "logger", log)

	// Create a new helloworld receiver.
	receiver := &Server{}

	// Create a request.
	request := &hw.SayHelloRequest{
		Name: "Tony Baloney",
	}

	// Call the SayHello method.
	response, err := receiver.SayHello(updatedCtx, request)
	expected := "Hello, " + request.Name

	// Assert that the error is nil.
	assert.Nil(t, err)
	assert.Equal(t, expected, response.Message)
}

// TestSayHello_Failure is confirming our HelloWorld-like returns an error when no name is included
func TestSayHello_Failure(t *testing.T) {
	log := zap.L()
	ctx := context.Background()
	updatedCtx := context.WithValue(ctx, "logger", log)

	// Create a new helloworld receiver.
	receiver := &Server{}

	// Create a request.
	request := &hw.SayHelloRequest{}

	expected := status.Error(codes.InvalidArgument, "name is required")
	// Call the SayHello method.
	response, err := receiver.SayHello(updatedCtx, request)

	// Assert that the error is nil.
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expected.Error())
}
