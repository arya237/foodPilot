package auth

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type AutoSaveRequest struct {
	AutoSave bool `json:"autosave" binding:"required"`
}

type AutoSaveResponse struct {
	Message string `json:"message"`
}