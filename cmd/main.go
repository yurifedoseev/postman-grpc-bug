package main

import (
	"context"
	"fmt"
	"github.com/yurifedoseev/postman-grpc-bug/pkg/app"
	"github.com/yurifedoseev/postman-grpc-bug/pkg/proto/proto/api"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const defaultPort = 8080

func run() error {
	slog.Info("start application")

	grpcServer := grpc.NewServer()
	service := &app.BackendService{}
	api.RegisterBackendServer(grpcServer, service) // add backend service
	reflection.Register(grpcServer)                // register server reflection

	port := defaultPort
	var err error
	if portRaw := os.Getenv("GRPC_PORT"); portRaw != "" {
		port, err = strconv.Atoi(portRaw)
		if err != nil {
			return fmt.Errorf("failed to part GRPC_PORT: %w", err)
		}
	}

	ctx := context.Background()
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		// start grpc server
		addr := fmt.Sprintf(":%d", port)
		slog.Info("start grpc server", "addr", addr)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}
		return grpcServer.Serve(listener)
	})

	// handle sigterm
	wg.Go(func() error {
		done := make(chan os.Signal, 1)
		signal.Notify(done, syscall.SIGTERM)
		<-done
		return fmt.Errorf("SIGTERM received")
	})

	return wg.Wait()
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
