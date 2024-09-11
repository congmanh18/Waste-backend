package handler

import "smart-waste/domain/wastebin/usecase"

type WasteBinHandler struct {
	CreateWasteBinUsecase *usecase.CreateWasteBinUsecase
	DeleteWasteBinUsecase *usecase.DeleteWasteBinUsecase
	UpdateWasteBinUsecase *usecase.UpdateWasteBinUsecase
	ReadWasteBinUsecase   *usecase.ReadWasteBinUsecase
}
