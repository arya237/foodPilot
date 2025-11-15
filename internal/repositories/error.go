package repositories

import "errors"

var (
	ErrorInvalidUID    = errors.New("user not found")
	ErrorInvalidUName  = errors.New("user not found")
	ErrorDuplicateUser = errors.New("duplicate user")
	ErrorNoUser        = errors.New("there is no user in database")
	ErrorInvalidFID    = errors.New("food not found")
	ErrorDuplicateFood = errors.New("duplicate food")
	ErrorNoFood        = errors.New("there is no food in database")
	ErrorNorate        = errors.New("there is no rate in database")
)
