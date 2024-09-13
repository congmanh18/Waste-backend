package usecase

import (
	"context"
	"smart-waste/domain/user/entity"
	"smart-waste/domain/user/repository"

	"gorm.io/gorm"
)

type UpdateUserUsecase struct {
	userRepo repository.UserRepo
}

func NewUpdateUserUsecase(db *gorm.DB) *UpdateUserUsecase {
	return &UpdateUserUsecase{
		userRepo: repository.NewUserRepo(db),
	}
}

func (u *UpdateUserUsecase) ExecuteUpdateUser(ctx context.Context, id string, user *entity.User) error {
	return u.userRepo.Update(ctx, &id, user)
}
