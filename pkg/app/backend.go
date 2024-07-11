package app

import (
	"context"
	"fmt"
	"github.com/yurifedoseev/postman-grpc-bug/pkg/proto/proto/api"
	"log/slog"
)

type BackendService struct {
}

func (b BackendService) Call(_ context.Context, request *api.CallRequest) (*api.CallResponse, error) {
	slog.Info("calling api method", "request", request)
	return &api.CallResponse{
		Message: fmt.Sprintf("response name %s with map %v", request.Name, request.Dict),
	}, nil
}
