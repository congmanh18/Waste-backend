package usecase

import (
	"context"
	"smart-waste/domain/report/entity"
	"smart-waste/domain/report/repository"

	"gorm.io/gorm"
)

type GetReportsByUserIDUsecase struct {
	reportRepo repository.ReportRepo
}

func NewGetReportsByUserIDUsecase(db *gorm.DB) *GetReportsByUserIDUsecase {
	return &GetReportsByUserIDUsecase{reportRepo: repository.NewReportRepo(db)}
}

func (u *GetReportsByUserIDUsecase) Execute(ctx context.Context, userID *string) (*[]entity.Report, error) {
	return u.reportRepo.GetAllByUserID(ctx, userID)
}
