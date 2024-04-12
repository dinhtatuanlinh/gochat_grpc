package domain

import (
	"context"
	"go_chat/internal/delivery/input"
	"go_chat/libraries/entities"
)

type UseCase interface {
	ShowNameUseCase(ctx context.Context) (result *entities.ShowNameUseCaseEntity, err error)
	Test(ctx context.Context, ipt *input.TestInput) (result *entities.TestUsecaseEntity, err error)
}
