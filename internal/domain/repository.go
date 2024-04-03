package domain

import (
	"context"
	"go_chat/libraries/entities"
)

type Repository interface {
	ShowNameRepo(ctx context.Context) (result *entities.ShowNameRepoEntity, err error)
}
