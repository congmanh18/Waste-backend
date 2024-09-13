package req

type UpdateUserReq struct {
	ID        string  `json:"id"`
	FirstName *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  *string `json:"last_name" validate:"required,min=2,max=100"`
	Gender    *string `json:"gender" validate:"required,eq=male|eq=female"`
	Category  *string `json:"category" validate:"required,eq=fulltime|eq=parttime"`
	Email     *string `json:"email"`
	Password  *string `json:"password" validate:"required,min=6"`
}
