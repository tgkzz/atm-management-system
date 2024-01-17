package models

import "errors"

var (
	ErrInvalidEmail         error = errors.New("invalid email")
	ErrInvalidPassword      error = errors.New("invalid password")
	ErrUsernameAlreadyTaken error = errors.New("username is already taken")
	ErrEmailAlreadyTaken    error = errors.New("email is already taken")
	ErrIncorrectPassword    error = errors.New("incorrect password")
	ErrNoUser               error = errors.New("no such user")
)
