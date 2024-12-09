package handler

import (
	"smart-waste/domain/report/usecase"
	userRepo "smart-waste/domain/user/repository"
	wasteBinRepo "smart-waste/domain/wastebin/repository"
)

type ReportHandler struct {
	CreateReportUsecase           *usecase.CreateReportUsecase
	DeleteReportUsecase           *usecase.DeleteReportUsecase
	GetAllReportsUsecase          *usecase.GetAllReportsUsecase
	GetReportByIDUsecase          *usecase.GetReportByIDUsecase
	GetReportsByDateUsecase       *usecase.GetReportsByDateUsecase
	GetReportsByUserIDUsecase     *usecase.GetReportsByUserIDUsecase
	GetLast                       *usecase.GetLatestByWasteBinID
	GetReportsByWasteBinIDUsecase *usecase.GetReportsByWasteBinIDUsecase
	WasteBinRepo                  wasteBinRepo.WasteBinRepo
	UserRepo                      userRepo.UserRepo
}

func NewReportHandler(uc *usecase.CreateReportUsecase, wasteBinRepo wasteBinRepo.WasteBinRepo, userRepo userRepo.UserRepo) *ReportHandler {
	return &ReportHandler{
		CreateReportUsecase: uc,
		WasteBinRepo:        wasteBinRepo,
		UserRepo:            userRepo,
	}
}
