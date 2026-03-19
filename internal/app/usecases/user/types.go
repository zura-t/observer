package userUsecase

import "time"

type RegisterUser struct {
	ID         uint64
	Email      string
	Password   string
	Name       string
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Login struct {
	Email    string
	Password string
}

type UserResponse struct {
	ID         uint64
	Email      string
	Name       string
	IsVerified bool
	CreatedAt  int64
	UpdatedAt  int64
}

type NewAccessToken struct {
	AccessToken string
}

type UpdateUser struct {
	ID         uint64
	Name       string
	Email      string
	IsVerified *bool
}

type ChangePassword struct {
	ID          uint64
	OldPassword string
	NewPassword string
}

type UserSearchFilter struct {
	ID         *uint64
	Email      string
	Name       string
	IsVerified *bool
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
