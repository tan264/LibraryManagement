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

type FilterUserRequest struct {
	UserID    uint   `form:"user_id"`
	CreatedAt string `form:"created_at"`
	Phone     string `form:"phone"`
}

type CreateBookRequest struct {
	Title     string `json:"title" binding:"required"`
	LibraryID uint   `json:"library_id" binding:"required"`
}

type FilterBookRequest struct {
	BookID    uint   `form:"book_id"`
	CreatedAt string `form:"created_at"`
}

type UpdateBookRequest struct {
	BookID    uint    `json:"book_id" binding:"required"`
	Title     *string `json:"title"`
	LibraryID uint    `json:"library_id"`
}
