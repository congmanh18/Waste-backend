package usecase

import (
	"context"
	"smart-waste/domain/report/repository"

	"gorm.io/gorm"
)

type DeleteReportUsecase struct {
	reportRepo repository.ReportRepo
}

func NewDeleteReportUsecase(db *gorm.DB) *DeleteReportUsecase {
	return &DeleteReportUsecase{reportRepo: repository.NewReportRepo(db)}
}

func (u *DeleteReportUsecase) Execute(ctx context.Context, id *string) error {
	return u.reportRepo.DeleteReport(ctx, id)
}
