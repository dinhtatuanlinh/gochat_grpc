package repository

import (
	"context"
	"go_chat/libraries/entities"
)

func (repo *repository) ShowNameRepo(ctx context.Context) (result *entities.ShowNameRepoEntity, err error) {

	return &entities.ShowNameRepoEntity{
		Name: "linh",
	}, nil
}
