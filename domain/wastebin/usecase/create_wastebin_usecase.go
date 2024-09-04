package usecase

import (
	"context"
	"smart-waste/domain/wastebin/entity"
	"smart-waste/domain/wastebin/repository"

	"gorm.io/gorm"
)

type CreateWasteBinUsecase struct {
	wasteBinRepo repository.WasteBinRepo
}

func NewCreateWasteBinUsecase(db *gorm.DB) *CreateWasteBinUsecase {
	return &CreateWasteBinUsecase{
		wasteBinRepo: repository.NewWasteBinRepo(db),
	}
}

func (c CreateWasteBinUsecase) ExecuteCreateWasteBin(ctx context.Context, wasteBinRepo *entity.WasteBin) error {
	return c.wasteBinRepo.Save(ctx, wasteBinRepo)
}
