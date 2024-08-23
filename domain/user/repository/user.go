package repository

import (
	"context"
	"smart-waste/domain/user/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	Save(ctx context.Context, user entity.User) error
	Update(ctx context.Context, id string, user entity.User) error
}

type userRepoImpl struct {
	gorm *gorm.DB
}

func NewUserRepo(gorm *gorm.DB) UserRepo {
	return &userRepoImpl{
		gorm: gorm,
	}
}

func (u *userRepoImpl) Save(ctx context.Context, user entity.User) error {
	return u.gorm.WithContext(ctx).Debug().Save(&user).Error
}

func (u *userRepoImpl) Update(ctx context.Context, id string, user entity.User) error {
	return u.gorm.WithContext(ctx).Debug().Where("id = ?", id).Updates(&user).Error
}
