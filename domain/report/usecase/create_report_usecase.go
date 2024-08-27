package usecase

import (
	"context"
	"smart-waste/domain/report/entity"
	"smart-waste/domain/report/repository"

	"gorm.io/gorm"
)

type CreateReportUsecase struct {
	reportRepo repository.ReportRepo
}

func NewCreateReportUsecase(db *gorm.DB) *CreateReportUsecase {
	return &CreateReportUsecase{
		reportRepo: repository.NewReportRepo(db),
	}
}

func (c *CreateReportUsecase) ExecuteCreateReport(ctx context.Context, report entity.Report) error {
	return c.reportRepo.Save(ctx, report)
}
