package delivery

import (
	"context"
	protocolBuffer "go_chat/internal/delivery/proto/protocol_buffer"
)

func (delivery *GrpcDelivery) ShowName(ctx context.Context, req *protocolBuffer.ShowNameRequest) (res *protocolBuffer.ShowNameResponse, err error) {
	useCaseResult, err := delivery.useCase.ShowNameUseCase(ctx)
	return &protocolBuffer.ShowNameResponse{Name: useCaseResult.Name}, err
}
