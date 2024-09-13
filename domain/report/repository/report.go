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

func (r *reportRepoImpl) GetByDate(ctx context.Context, date *string) (*[]entity.Report, error) {
	var reports []entity.Report
	err := r.gorm.Where("date =?", date).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return &reports, nil
}

func (r *reportRepoImpl) GetAll(ctx context.Context) (*[]entity.Report, error) {
	var reports []entity.Report
	err := r.gorm.Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return &reports, nil
}
