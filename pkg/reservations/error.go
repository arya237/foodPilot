package reservations

import "errors"

var (
	ErrorInvalidToken = errors.New("invalid token")
	ErrorInternal     = errors.New("internal error")
)
