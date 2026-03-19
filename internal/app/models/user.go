package models

import "time"

type UserDB struct {
	ID         uint64
	Email      string
	Password   string
	Name       string
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type User struct {
	ID         uint64
	Name       string
	Email      string
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserWithToken struct {
	User   *User
	AccessToken  string
	AccessTokenExpiresAt time.Time
	RefreshToken  string
	RefreshTokenExpiresAt time.Time
}

type NullString struct {
	Value string
	IsNil bool
}

type NullBool struct {
	Value bool
	IsNil bool
}
