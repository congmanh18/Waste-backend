package req

type LoginUserReq struct {
	Phone    string  `json:"phone" validate:"required"`
	Password *string `json:"password"`
}
