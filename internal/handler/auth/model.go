package auth

// LoginRequest is the request body for user login
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"your username"`
	Password string `json:"password" binding:"required" example:"your password"`
}

// LoginResponse is the response returned after successful login
type LoginResponse struct {
	Token string `json:"token" example:"generated token"`
}

// ErrorResponse is returned when an error occurs
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}