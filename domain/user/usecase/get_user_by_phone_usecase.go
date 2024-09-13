package usecase

import (
	"context"
	"smart-waste/domain/user/entity"
	"smart-waste/domain/user/repository"

	"gorm.io/gorm"
)

type GetUserByPhoneUsecase struct {
	userRepo repository.UserRepo
}

func NewGetUserByPhoneUsecase(db *gorm.DB) *GetUserByPhoneUsecase {
	return &GetUserByPhoneUsecase{
		userRepo: repository.NewUserRepo(db),
	}
}

func (f *GetUserByPhoneUsecase) ExecuteGetUserByPhone(ctx context.Context, phone *string) (*entity.User, error) {
	return f.userRepo.GetByPhone(ctx, phone)
}
