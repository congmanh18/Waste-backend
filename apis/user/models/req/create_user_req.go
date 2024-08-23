package req

type CreateUserReq struct {
	ID        string  `json:"id" gorm:"primaryKey"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Gender    *string `json:"gender"`
	Role      *string `json:"role"`
	Category  *string `json:"category"`
	Email     *string `json:"email"`
	Phone     string  `json:"phone" validate:"required"`
	Password  *string `json:"password"`
}
