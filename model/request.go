package model

type RegisterUserRequest struct {
	Username  string  `json:"username" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EditAccountRequest struct {
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}
