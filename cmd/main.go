package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"go_chat/internal/delivery"
	protocolBuffer "go_chat/internal/delivery/proto/protocol_buffer"
	"go_chat/internal/repository"
	"go_chat/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	//ctx := context.Background()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	Message := make(chan *protocolBuffer.Message)
	defer close(Message)
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			selector.UnaryServerInterceptor(
				auth.UnaryServerInterceptor(authenticator),
				selector.MatchFunc(authMatcher),
			),
			selector.UnaryServerInterceptor(
				auth.UnaryServerInterceptor(authenticatortwo),
				selector.MatchFunc(authMatcher),
			),
		),
	)

	repo := repository.NewRepository()
	uc := usecase.NewUsecase(repo)
	delivery.NewGrpcDelivery(s, uc)

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("bbbbbb")
}

func authenticatortwo(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	// TODO: This is example only, perform proper Oauth/OIDC verification!
	if token != "yolo" {
		return nil, status.Error(codes.Unauthenticated, "invalid auth token")
	}
	// NOTE: You can also pass the token in the context for further interceptors or gRPC service code.
	return ctx, nil
}

func authenticator(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	// TODO: This is example only, perform proper Oauth/OIDC verification!
	if token != "yolo" {
		return nil, status.Error(codes.Unauthenticated, "invalid auth token")
	}
	// NOTE: You can also pass the token in the context for further interceptors or gRPC service code.
	return ctx, nil
}

func authMatcher(ctx context.Context, callMeta interceptors.CallMeta) bool {
	fmt.Println(callMeta)
	fmt.Println(healthpb.Health_ServiceDesc.ServiceName)
	fmt.Println(callMeta.Service)
	//return true
	//return healthpb.Health_ServiceDesc.ServiceName != callMeta.Service
	return callMeta.Service == "ShowName"
}
