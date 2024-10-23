package repository

import (
	"context"
	"smart-waste/domain/report/entity"

	"gorm.io/gorm"
)

// Định nghĩa giao diện ReportRepo cho các thao tác với Report
type ReportRepo interface {
	Save(ctx context.Context, report *entity.Report) error
	GetByID(ctx context.Context, id *string) (*entity.Report, error)
	GetByDate(ctx context.Context, date *string) (*[]entity.Report, error)
	GetAll(ctx context.Context) (*[]entity.Report, error)
	GetAllByUserID(ctx context.Context, userID *string) (*[]entity.Report, error)
	GetAllByWasteBinID(ctx context.Context, wasteBinID *string) (*[]entity.Report, error)
	DeleteReport(ctx context.Context, id *string) error
}

// Cấu trúc chứa đối tượng GORM DB để triển khai ReportRepo
type reportRepoImpl struct {
	db *gorm.DB
}

// Hàm khởi tạo Repository
func NewReportRepo(db *gorm.DB) ReportRepo {
	return &reportRepoImpl{db: db}
}

// Lưu báo cáo mới
func (r *reportRepoImpl) Save(ctx context.Context, report *entity.Report) error {
	return r.db.WithContext(ctx).Create(&report).Error
}

// Lấy báo cáo theo ID
func (r *reportRepoImpl) GetByID(ctx context.Context, id *string) (*entity.Report, error) {
	var report entity.Report
	if err := r.db.WithContext(ctx).First(&report, "id = ?", *id).Error; err != nil {
		return nil, err
	}
	return &report, nil
}

// Lấy báo cáo theo ngày
func (r *reportRepoImpl) GetByDate(ctx context.Context, date *string) (*[]entity.Report, error) {
	var reports []entity.Report
	err := r.db.WithContext(ctx).Debug().
		Where("DATE(created_at) = ?", *date).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return &reports, nil
}

// Lấy tất cả báo cáo
func (r *reportRepoImpl) GetAll(ctx context.Context) (*[]entity.Report, error) {
	var reports []entity.Report
	if err := r.db.WithContext(ctx).Find(&reports).Error; err != nil {
		return nil, err
	}
	return &reports, nil
}

// Lấy báo cáo theo User ID
func (r *reportRepoImpl) GetAllByUserID(ctx context.Context, userID *string) (*[]entity.Report, error) {
	var reportList []entity.Report
	err := r.db.WithContext(ctx).Debug().
		Where("user_id = ?", *userID).Find(&reportList).Error
	if err != nil {
		return nil, err
	}
	return &reportList, nil
}

// Lấy báo cáo theo WasteBin ID
func (r *reportRepoImpl) GetAllByWasteBinID(ctx context.Context, wasteBinID *string) (*[]entity.Report, error) {
	var reportList []entity.Report
	err := r.db.WithContext(ctx).Debug().
		Where("wastebin_id = ?", *wasteBinID).Find(&reportList).Error
	if err != nil {
		return nil, err
	}
	return &reportList, nil
}

// Xóa báo cáo theo ID
func (r *reportRepoImpl) DeleteReport(ctx context.Context, id *string) error {
	return r.db.WithContext(ctx).Delete(&entity.Report{}, "id = ?", *id).Error
}
