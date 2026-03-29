package models

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")

	ErrUserNotFound       = errors.New("user not found")
	ErrDiaryEntryNotFound = errors.New("diary entry not found")
	ErrHabitNotFound      = errors.New("habit not found")
	ErrHabitLogNotFound   = errors.New("habit log not found")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidEmail        = errors.New("invalid email address")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrAccountNotVerified  = errors.New("account not verified")
	ErrInvalidToken        = errors.New("invalid token")
	ErrTokenExpired        = errors.New("token expired")
	ErrVerificationExpired = errors.New("verification expired")
	ErrRateLimitExceeded   = errors.New("rate limit exceeded")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
)
