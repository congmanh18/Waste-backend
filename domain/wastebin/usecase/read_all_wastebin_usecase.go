package usecase

import (
	"context"
	"smart-waste/domain/wastebin/entity"
	"smart-waste/domain/wastebin/repository"

	"gorm.io/gorm"
)

type ReadAllWasteBinUsecase struct {
	wasteBinRepo repository.WasteBinRepo
}

func NewReadAllWasteBinUsecase(db *gorm.DB) *ReadAllWasteBinUsecase {
	return &ReadAllWasteBinUsecase{
		wasteBinRepo: repository.NewWasteBinRepo(db),
	}
}

func (r *ReadAllWasteBinUsecase) ReadAllWasteBins(ctx context.Context) ([]entity.WasteBin, error) {
	wasteBins, err := r.wasteBinRepo.FindAll(ctx)
	if err != nil {
		return nil, err // Trả về lỗi nếu có
	}
	return wasteBins, nil
}
