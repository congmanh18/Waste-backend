package repository

import (
	"context"
	"smart-waste/domain/report/entity"

	"gorm.io/gorm"
)

type ReportRepo interface {
	Save(ctx context.Context, report entity.Report) error
	Delete(ctx context.Context, id string) error
}

type ReportRepoImpl struct {
	gorm *gorm.DB
}

func NewReportRepo(db *gorm.DB) *ReportRepoImpl {
	return &ReportRepoImpl{gorm: db}
}

func (r *ReportRepoImpl) Save(ctx context.Context, report entity.Report) error {
	return r.gorm.Create(&report).Error
}

func (r *ReportRepoImpl) Delete(ctx context.Context, id string) error {
	return r.gorm.Delete(&entity.Report{}, id).Error
}

// ... Other methods for interacting with reports, if needed ...
