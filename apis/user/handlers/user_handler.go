package handler

import usecase "smart-waste/domain/user/usecase"

type UserHandler struct {
	CreateUserUsecase   *usecase.CreateUserUsecase
	UpdateUserUsecase   *usecase.UpdateUserUsecase
	DeleteUserUsecase   *usecase.DeleteUserUsecase
	FindUserByIDUsecase *usecase.FindUserByIDUsecase
	FindAllUserUsecase  *usecase.FindAllUserUsecase
}
