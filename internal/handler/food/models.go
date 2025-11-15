package food





// ErrorResponse is returned when an error occurs
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// MessageResponse is returned when successful and message occurs
type MessageResponse struct {
	Message string `json:"message" example:"message"`
}
