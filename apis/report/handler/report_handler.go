package handler

import "smart-waste/domain/report/usecase"

type ReportHandler struct {
	CreateReportUsecase usecase.CreateReportUsecase
	DeleteReportUsecase usecase.DeleteReportUsecase
}
