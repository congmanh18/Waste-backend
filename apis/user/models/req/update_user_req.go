package req

type UpdateUserReq struct {
	ID        string  `json:"id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Gender    *string `json:"gender"`
	Category  *string `json:"category"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
}
