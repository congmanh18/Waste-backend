package usecase

import (
	"context"
	"smart-waste/domain/report/entity"
	"smart-waste/domain/report/repository"

	"gorm.io/gorm"
)

// Định nghĩa cấu trúc GetReportsByDateUsecase
type GetReportsByDateUsecase struct {
	reportRepo repository.ReportRepo
}

// Hàm khởi tạo GetReportsByDateUsecase
func NewGetReportsByDateUsecase(db *gorm.DB) *GetReportsByDateUsecase {
	return &GetReportsByDateUsecase{reportRepo: repository.NewReportRepo(db)}
}

// Execute chạy logic để lấy báo cáo theo ngày
func (u *GetReportsByDateUsecase) Execute(ctx context.Context, date *string) (*[]entity.Report, error) {
	return u.reportRepo.GetByDate(ctx, date)
}
