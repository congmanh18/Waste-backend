package usecase

import (
	"context"
	"smart-waste/domain/wastebin/entity"
	"smart-waste/domain/wastebin/repository"

	"gorm.io/gorm"
)

type ReadWasteBinUsecase struct {
	wasteBinRepo repository.WasteBinRepo
}

func NewReadWasteBinUsecase(db *gorm.DB) *ReadWasteBinUsecase {
	return &ReadWasteBinUsecase{
		wasteBinRepo: repository.NewWasteBinRepo(db),
	}
}

// ReadWasteBinByID reads a waste bin by its ID.
func (r *ReadWasteBinUsecase) ReadWasteBinByID(ctx context.Context, id string) (*entity.WasteBin, error) {
	return r.wasteBinRepo.FindById(ctx, &id)
}
