package handler

import usecase "smart-waste/domain/user/usecase"

type UserHandler struct {
	CreateUserUsecase     *usecase.CreateUserUsecase
	GetUserByPhoneUsecase *usecase.GetUserByPhoneUsecase
	UpdateUserUsecase     *usecase.UpdateUserUsecase
	DeleteUserUsecase     *usecase.DeleteUserUsecase
	FindUserByIDUsecase   *usecase.FindUserByIDUsecase
	FindAllUserUsecase    *usecase.FindAllUserUsecase
	// AdminOnlyHandler      *usecase.FindAllUserUsecase
}
