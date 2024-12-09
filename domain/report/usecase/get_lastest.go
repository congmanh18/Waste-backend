package usecase

import (
	"context"
	"smart-waste/domain/report/entity"
	"smart-waste/domain/report/repository"

	"gorm.io/gorm"
)

type GetLatestByWasteBinID struct {
	reportRepo repository.ReportRepo
}

func NewGetLatestByWasteBinID(db *gorm.DB) *GetLatestByWasteBinID {
	return &GetLatestByWasteBinID{reportRepo: repository.NewReportRepo(db)}
}

func (c *GetLatestByWasteBinID) Execute(ctx context.Context, wasteBinID *string) (*entity.Report, error) {
	report, err := c.reportRepo.GetLatestByWasteBinID(ctx, wasteBinID)
	if err != nil {
		return nil, err
	}

	return report, nil

}
