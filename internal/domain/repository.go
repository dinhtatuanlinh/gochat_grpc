package domain

import (
	"context"
	"go_chat/internal/delivery/input"
	"go_chat/libraries/entities"
)

type Repository interface {
	ShowNameRepo(ctx context.Context) (result *entities.ShowNameRepoEntity, err error)
	Test(ctx context.Context, ipt *input.TestInput) (result *entities.TestRepoEntity, err error)
}
