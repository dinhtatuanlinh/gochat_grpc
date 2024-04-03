package main

import (
	"go_chat/internal/delivery"
	protocolBuffer "go_chat/internal/delivery/proto/protocol_buffer"
	"go_chat/internal/repository"
	"go_chat/internal/usecase"
	"google.golang.org/grpc"
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
	s := grpc.NewServer()

	repo := repository.NewRepository()
	uc := usecase.NewUsecase(repo)
	delivery.NewGrpcDelivery(s, uc)

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
