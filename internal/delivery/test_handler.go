package delivery

import (
	"context"
	"go_chat/internal/delivery/input"
	protocolBuffer "go_chat/internal/delivery/proto/protocol_buffer"
)

func (delivery *GrpcDelivery) Test(ctx context.Context, req *protocolBuffer.Request) (res *protocolBuffer.Response, err error) {
	ipt := input.TestInput{
		Data: req.Data,
	}
	useCaseResult, err := delivery.useCase.Test(ctx, &ipt)
	return &protocolBuffer.Response{Result: useCaseResult.Data}, err
}
