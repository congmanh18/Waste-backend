package usecase

import (
	"context"
	"smart-waste/domain/report/repository"

	"gorm.io/gorm"
)

type DeleteReportUsecase struct {
	reportRepo repository.ReportRepo
}

func NewDeleteReportUsecase(db *gorm.DB) DeleteReportUsecase {
	return DeleteReportUsecase{
		reportRepo: repository.NewReportRepo(db),
	}
}

func (u *DeleteReportUsecase) ExecuteDeleteReport(ctx context.Context, id string) error {
	return u.reportRepo.Delete(ctx, &id)
}
