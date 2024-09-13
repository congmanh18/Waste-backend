package req

type CreateUserReq struct {
	ID        string  `json:"id"`
	FirstName *string `json:"first_name" validate:"required,min=2,max=100"`
	LastName  *string `json:"last_name" validate:"required,min=2,max=100"`
	Gender    *string `json:"gender" validate:"required,eq=male|eq=female"`
	Role      *string `json:"role" validate:"required,eq=admin|eq=staff"`
	Category  *string `json:"category" validate:"required,eq=fulltime|eq=parttime"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone" validate:"required"`
	Password  *string `json:"password" validate:"required,min=6"`
}
