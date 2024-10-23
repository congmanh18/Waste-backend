package usecase

import (
	"context"
	"smart-waste/domain/report/entity"
	"smart-waste/domain/report/repository"

	"gorm.io/gorm"
)

// Định nghĩa cấu trúc GetAllReportsUsecase
type GetAllReportsUsecase struct {
	reportRepo repository.ReportRepo
}

// Hàm khởi tạo GetAllReportsUsecase
func NewGetAllReportsUsecase(db *gorm.DB) *GetAllReportsUsecase {
	return &GetAllReportsUsecase{reportRepo: repository.NewReportRepo(db)}
}

// Execute chạy logic để lấy tất cả các báo cáo
func (u *GetAllReportsUsecase) Execute(ctx context.Context) (*[]entity.Report, error) {
	return u.reportRepo.GetAll(ctx)
}
