package usecase

import (
	"context"
	"go_chat/libraries/entities"
)

func (uc *useCase) ShowNameUseCase(ctx context.Context) (result *entities.ShowNameUseCaseEntity, err error) {
	repoResult, err := uc.repo.ShowNameRepo(ctx)
	return &entities.ShowNameUseCaseEntity{
		Name: repoResult.Name,
	}, err
}
