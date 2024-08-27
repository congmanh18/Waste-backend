package usecase

import (
	"context"
	"smart-waste/domain/user/entity"
	"smart-waste/domain/user/repository"

	"gorm.io/gorm"
)

type FindUserByIDUsecase struct {
	userRepo repository.UserRepo
}

func NewFindUserByIDUsecase(db *gorm.DB) *FindUserByIDUsecase {
	return &FindUserByIDUsecase{
		userRepo: repository.NewUserRepo(db),
	}
}

func (f *FindUserByIDUsecase) ExecuteFindById(ctx context.Context, id string) (*entity.User, error) {
	return f.userRepo.FindById(ctx, id)
}
