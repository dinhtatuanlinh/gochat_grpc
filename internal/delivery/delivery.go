package delivery

import (
	protocolBuffer "go_chat/internal/delivery/proto/protocol_buffer"
	"go_chat/internal/domain"
	"google.golang.org/grpc"
)

type ClientMsg struct {
	sender string
	conn   protocolBuffer.Services_SendMessageToChannelServer
}

type GrpcDelivery struct {
	protocolBuffer.UnimplementedServicesServer
	useCase           domain.UseCase
	message           chan *protocolBuffer.Message
	connections       map[string]int
	clientConnections map[string][]ClientMsg
}

func NewGrpcDelivery(grpcServer *grpc.Server, uc domain.UseCase) {

	protocolBuffer.RegisterServicesServer(grpcServer, &GrpcDelivery{
		useCase:           uc,
		message:           make(chan *protocolBuffer.Message),
		connections:       make(map[string]int),
		clientConnections: make(map[string][]ClientMsg),
	})
}
