package repository

import (
	"context"
	"smart-waste/domain/user/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	Save(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, id string, user *entity.User) error
	FindById(ctx context.Context, id string) (*entity.User, error)
	FindAll(ctx context.Context) (*[]entity.User, error)
	Delete(ctx context.Context, id string) error
}

type userRepoImpl struct {
	gorm *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepoImpl{
		gorm: db,
	}
}

func (u *userRepoImpl) Save(ctx context.Context, user *entity.User) error {
	return u.gorm.WithContext(ctx).Debug().Save(&user).Error
}

func (u *userRepoImpl) Update(ctx context.Context, id string, user *entity.User) error {
	return u.gorm.WithContext(ctx).Debug().Where("id = ?", id).Updates(&user).Error
}

func (u *userRepoImpl) FindById(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	if err := u.gorm.WithContext(ctx).Debug().Where("id = ?", id).First(&user).Error; err != nil { //+
		return nil, err
	}
	return &user, nil
}

func (u *userRepoImpl) FindAll(ctx context.Context) (*[]entity.User, error) {
	var userList []entity.User
	if err := u.gorm.WithContext(ctx).Debug().Find(&userList).Error; err != nil { //+
		return nil, err
	}
	return &userList, nil
}

func (u *userRepoImpl) Delete(ctx context.Context, id string) error {
	return u.gorm.WithContext(ctx).Debug().Delete(&entity.User{}, id).Error
}
