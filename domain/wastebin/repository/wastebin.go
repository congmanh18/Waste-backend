package repository

import (
	"context"
	"smart-waste/domain/wastebin/entity"

	"gorm.io/gorm"
)

type WasteBinRepo interface {
	Save(ctx context.Context, wastebin *entity.WasteBin) error              //
	Update(ctx context.Context, id string, wastebin *entity.WasteBin) error //
	FindById(ctx context.Context, id string) (*entity.WasteBin, error)      //
	FindAll(ctx context.Context) (*[]entity.WasteBin, error)                //
	Delete(ctx context.Context, id string) error                            //
}

type wasteBinRepoImpl struct {
	gorm *gorm.DB
}

func NewWasteBinRepo(db *gorm.DB) WasteBinRepo {
	return &wasteBinRepoImpl{
		gorm: db,
	}
}

func (w *wasteBinRepoImpl) Save(ctx context.Context, wastebin *entity.WasteBin) error {
	return w.gorm.WithContext(ctx).Debug().Save(&wastebin).Error
}

func (w *wasteBinRepoImpl) Update(ctx context.Context, id string, wastebin *entity.WasteBin) error {
	return w.gorm.WithContext(ctx).Debug().Where("id = ?", id).Updates(wastebin).Error
}

func (w *wasteBinRepoImpl) FindById(ctx context.Context, id string) (*entity.WasteBin, error) {
	var wastebin entity.WasteBin
	if err := w.gorm.WithContext(ctx).Debug().Where("id = ?", id).First(&wastebin).Error; err != nil {
		return nil, err
	}
	return &wastebin, nil
}

func (w *wasteBinRepoImpl) FindAll(ctx context.Context) (*[]entity.WasteBin, error) {
	var wasteBinList []entity.WasteBin
	if err := w.gorm.WithContext(ctx).Debug().Find(&wasteBinList).Error; err != nil {
		return nil, err
	}
	return &wasteBinList, nil
}

func (w *wasteBinRepoImpl) Delete(ctx context.Context, id string) error {
	return w.gorm.WithContext(ctx).Debug().Delete(&entity.WasteBin{}, id).Error
}
