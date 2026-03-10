package controller

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GetUserByIDRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type GetUserByEmailRequest struct {
	Email string `form:"email" binding:"required,email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
}

type DeleteUserRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}
