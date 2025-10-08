package user

// AutoSaveRequest is the request body for enabling or disabling autosave
type AutoSaveRequest struct {
	AutoSave *bool `json:"autosave" binding:"required"`
}

// AutoSaveResponse is the response returned after updating the autosave setting
type AutoSaveResponse struct {
	Message string `json:"message"`
}

// RatesResponse is the response containing food ratings
type RatesResponse struct {
	Rates map[string]int `json:"rates" example:"foodName1:93,foodName2:74,foodName3:80"`
}

// ErrorResponse is returned when an error occurs
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}
