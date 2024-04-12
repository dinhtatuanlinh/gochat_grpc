package repository

import (
	"context"
	"fmt"
	"go_chat/internal/delivery/input"
	"go_chat/libraries/entities"
)

func (repo *repository) Test(ctx context.Context, ipt *input.TestInput) (result *entities.TestRepoEntity, err error) {
	return &entities.TestRepoEntity{
		Data: fmt.Sprintf("%s abc", ipt.Data),
	}, nil
}
