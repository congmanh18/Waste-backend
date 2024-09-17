package repository

import (
	"context"
	"smart-waste/domain/report/entity"

	"gorm.io/gorm"
)

type ReportRepo interface {
	Save(ctx context.Context, report *entity.Report) error
	GetByID(ctx context.Context, id *string) error
	GetByDate(ctx context.Context, date *string) (*[]entity.Report, error)
	GetAll(ctx context.Context) (*[]entity.Report, error)
	Delete(ctx context.Context, id *string) error
}

type reportRepoImpl struct {
	gorm *gorm.DB
}

func NewReportRepo(db *gorm.DB) *reportRepoImpl {
	return &reportRepoImpl{gorm: db}
}

func (r *reportRepoImpl) Save(ctx context.Context, report *entity.Report) error {
	return r.gorm.Create(&report).Error
}

func (r *reportRepoImpl) Delete(ctx context.Context, id *string) error {
	return r.gorm.Delete(&entity.Report{}, id).Error
}

func (r *reportRepoImpl) GetByID(ctx context.Context, id *string) error {
	return r.gorm.First(&entity.Report{}, id).Error
}

func (r *reportRepoImpl) GetAllByUserID(ctx context.Context, userID *string) (*[]entity.Report, error) {
	var reportList []entity.Report
	if err := r.gorm.WithContext(ctx).Debug().Where("user_id = ?", userID).Find(&reportList).Error; err != nil {
		return nil, err
	}
	return &reportList, nil
}

func (r *reportRepoImpl) GetAllByWasteBinID(ctx context.Context, wasteBinID *string) (*[]entity.Report, error) {
	var reportList []entity.Report
	if err := r.gorm.WithContext(ctx).Debug().Where("wastebin_id = ?", wasteBinID).Find(&reportList).Error; err != nil {
		return nil, err
	}
	return &reportList, nil
}

func (r *reportRepoImpl) GetByDate(ctx context.Context, date *string) (*[]entity.Report, error) {
	var reports []entity.Report
	err := r.gorm.WithContext(ctx).Debug().Where("DATE(created_at) = ?", date).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return &reports, nil
}

func (r *reportRepoImpl) GetAll(ctx context.Context) (*[]entity.Report, error) {
	var reports []entity.Report
	err := r.gorm.WithContext(ctx).Debug().Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return &reports, nil
}
