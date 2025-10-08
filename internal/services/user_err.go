package services

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrUserRegistration   = errors.New("user registration failed")
	ErrTokenGeneration    = errors.New("token generation failed")
	ErrUserNotRegistered  = errors.New("user not registered, please sign up first")
)
