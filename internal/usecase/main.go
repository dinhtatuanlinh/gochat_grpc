package usecase

import "go_chat/internal/domain"

type useCase struct {
	repo domain.Repository
}

func NewUsecase(repo domain.Repository) domain.UseCase {
	return &useCase{
		repo: repo,
	}
}
