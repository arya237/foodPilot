package repositories

import "errors"

var (
	ErrorInvalidID = errors.New("user not found")
	ErrorDuplicateUser  = errors.New("duplicate user")
)
