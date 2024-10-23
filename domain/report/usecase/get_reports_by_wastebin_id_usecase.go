package usecase

import (
	"context"
	"smart-waste/domain/report/entity"
	"smart-waste/domain/report/repository"

	"gorm.io/gorm"
)

type GetReportsByWasteBinIDUsecase struct {
	reportRepo repository.ReportRepo
}

func NewGetReportsByWasteBinIDUsecase(db *gorm.DB) *GetReportsByWasteBinIDUsecase {
	return &GetReportsByWasteBinIDUsecase{reportRepo: repository.NewReportRepo(db)}
}

func (u *GetReportsByWasteBinIDUsecase) Execute(ctx context.Context, wasteBinID *string) (*[]entity.Report, error) {
	return u.reportRepo.GetAllByWasteBinID(ctx, wasteBinID)
}
