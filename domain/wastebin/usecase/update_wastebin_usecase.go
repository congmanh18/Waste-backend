package usecase

import (
	"context"
	"smart-waste/domain/wastebin/entity"
	"smart-waste/domain/wastebin/repository"

	"gorm.io/gorm"
)

type UpdateWasteBinUsecase struct {
	wasteBinRepo repository.WasteBinRepo
}

func NewUpdateWasteBinUsecase(db *gorm.DB) *UpdateWasteBinUsecase {
	return &UpdateWasteBinUsecase{
		wasteBinRepo: repository.NewWasteBinRepo(db),
	}
}

// Update updates a waste bin with the provided ID.
func (uc *UpdateWasteBinUsecase) ExecuteUpdateWasteBin(ctx context.Context, id string, wasteBin *entity.WasteBin) error {
	return uc.wasteBinRepo.Update(ctx, &id, wasteBin)
}
