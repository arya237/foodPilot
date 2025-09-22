package repositories

import "errors"

var (
	ErrorInvalidID = errors.New("user not found")
	ErrorDuplicateUser  = errors.New("duplicate user")
	ErrorNoUser = errors.New("there is no user in database")
)
