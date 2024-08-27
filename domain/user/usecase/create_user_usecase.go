package usecase

import (
	"context"
	"smart-waste/domain/user/entity"
	"smart-waste/domain/user/repository"

	"gorm.io/gorm"
)

type CreateUserUsecase struct {
	userRepo repository.UserRepo
}

func NewCreateUserUsecase(db *gorm.DB) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepo: repository.NewUserRepo(db),
	}
}

func (c CreateUserUsecase) ExecuteCreateUser(ctx context.Context, user *entity.User) error {
	if err := user.IsValidUser(); err != nil {
		return err
	}
	return c.userRepo.Save(ctx, user)
}
