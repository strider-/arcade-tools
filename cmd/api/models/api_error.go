package models

type ApiError struct {
	Error string `json:"error"`
}

func NewApiError(e error) ApiError {
	return ApiError{Error: e.Error()}
}
