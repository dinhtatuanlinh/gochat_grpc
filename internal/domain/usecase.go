package domain

import (
	"context"
	"go_chat/libraries/entities"
)

type UseCase interface {
	ShowNameUseCase(ctx context.Context) (result *entities.ShowNameUseCaseEntity, err error)
}
