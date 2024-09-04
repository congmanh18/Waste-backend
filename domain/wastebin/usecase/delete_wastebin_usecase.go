package usecase

import (
	"context"
	"smart-waste/domain/wastebin/repository"

	"gorm.io/gorm"
)

type DeleteWasteBinUsecase struct {
	wasteBinRepo repository.WasteBinRepo
}

func NewDeleteUserUsecase(db *gorm.DB) *DeleteWasteBinUsecase {
	return &DeleteWasteBinUsecase{
		wasteBinRepo: repository.NewWasteBinRepo(db),
	}
}

func (d *DeleteWasteBinUsecase) ExecuteDeleteWasteBin(ctx context.Context, id string) error {
	return d.wasteBinRepo.Delete(ctx, id)
}
