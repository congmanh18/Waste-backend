package usecase

import (
	"context"
	"smart-waste/domain/user/repository"

	"gorm.io/gorm"
)

type DeleteUserUsecase struct {
	userRepo repository.UserRepo
}

func NewDeleteUserUsecase(db *gorm.DB) *DeleteUserUsecase {
	return &DeleteUserUsecase{
		userRepo: repository.NewUserRepo(db),
	}
}

func (u *DeleteUserUsecase) ExecuteDeleteUser(ctx context.Context, id string) error {
	return u.userRepo.Delete(ctx, id)
}
