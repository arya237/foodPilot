package user

type AutoSaveRequest struct {
	AutoSave *bool `json:"autosave" binding:"required"`
}

type AutoSaveResponse struct {
	Message string `json:"message"`
}
