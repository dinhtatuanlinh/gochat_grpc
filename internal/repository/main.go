package repository

import "go_chat/internal/domain"

type repository struct {
}

func NewRepository() domain.Repository {
	return &repository{}
}
