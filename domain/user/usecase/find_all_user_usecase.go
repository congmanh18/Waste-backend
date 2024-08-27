package usecase

import (
	"context"
	"smart-waste/domain/user/entity"
	"smart-waste/domain/user/repository"

	"gorm.io/gorm"
)

type FindAllUserUsecase struct {
	userRepo repository.UserRepo
}

func NewFindAllUserUsecase(db *gorm.DB) *FindAllUserUsecase {
	return &FindAllUserUsecase{
		userRepo: repository.NewUserRepo(db),
	}
}

func (f *FindAllUserUsecase) ExecuteFindAll(ctx context.Context) (*[]entity.User, error) {
	return f.userRepo.FindAll(ctx)
}
