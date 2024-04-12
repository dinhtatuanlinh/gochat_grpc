package usecase

import (
	"context"
	"go_chat/internal/delivery/input"
	"go_chat/libraries/entities"
)

func (uc *useCase) Test(ctx context.Context, req *input.TestInput) (result *entities.TestUsecaseEntity, err error) {
	repoResult, err := uc.repo.Test(ctx, req)
	return &entities.TestUsecaseEntity{
		Data: repoResult.Data,
	}, err
}
