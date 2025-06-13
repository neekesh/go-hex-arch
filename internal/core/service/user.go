package service

import (
	"context"
	"fmt"

	adapter_entities "github.com/thapakazi/go-hex-arch/internal/adapter/entities"
	"github.com/thapakazi/go-hex-arch/internal/adapter/storage/postgres/repository"
	"github.com/thapakazi/go-hex-arch/internal/core/entities"
	"github.com/thapakazi/go-hex-arch/internal/ports"
)

type UserService struct {
	userRepo ports.UserRepository
}

func NewUserService() *UserService {
	fmt.Println("NewUserService **** called")
	return &UserService{userRepo: repository.NewUserRepository()}
}

func (s *UserService) CreateUser(ctx context.Context, user *adapter_entities.User) error {
	userEntity := &entities.User{
		User: *user,
	}
	fmt.Println("UserService CreateUser ****", userEntity)
	return s.userRepo.CreateUser(ctx, userEntity)
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*entities.User, error) {
	return s.userRepo.GetUser(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *entities.User) error {
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.userRepo.DeleteUser(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context, params adapter_entities.QueryParams) ([]entities.User, error) {
	return s.userRepo.GetAllUsers(ctx, params)
}
