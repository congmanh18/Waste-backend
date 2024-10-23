package usecase

import (
	"context"
	"smart-waste/domain/report/entity"
	"smart-waste/domain/report/repository"

	"gorm.io/gorm"
)

// Định nghĩa cấu trúc GetReportByIDUsecase
type GetReportByIDUsecase struct {
	reportRepo repository.ReportRepo
}

// Hàm khởi tạo GetReportByIDUsecase
func NewGetReportByIDUsecase(db *gorm.DB) *GetReportByIDUsecase {
	return &GetReportByIDUsecase{reportRepo: repository.NewReportRepo(db)}
}

// Thực thi logic lấy báo cáo theo ID
func (u *GetReportByIDUsecase) Execute(ctx context.Context, id *string) (*entity.Report, error) {
	report, err := u.reportRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return report, nil
}
