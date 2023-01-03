package storage

import (
	"context"
	"github.com/shakhboznorbekov/token/user_api/models"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (string, error)
	GetByPKey(ctx context.Context, req *models.UserPrimarKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	// Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	// Delete(ctx context.Context, req *models.UserPrimarKey) error
}
