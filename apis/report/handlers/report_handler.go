package handler

import "smart-waste/domain/report/usecase"

type ReportHandler struct {
	CreateReportUsecase           *usecase.CreateReportUsecase
	DeleteReportUsecase           *usecase.DeleteReportUsecase
	GetAllReportsUsecase          *usecase.GetAllReportsUsecase
	GetReportByIDUsecase          *usecase.GetReportByIDUsecase
	GetReportsByDateUsecase       *usecase.GetReportsByDateUsecase
	GetReportsByUserIDUsecase     *usecase.GetReportsByUserIDUsecase
	GetReportsByWasteBinIDUsecase *usecase.GetReportsByWasteBinIDUsecase
}
