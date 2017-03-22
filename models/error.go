package models

// ErrorResponse is used to respond that an error has occurred
type ErrorResponse struct {
	Status  int    `json:"status"`
	Title   string `json:"title"`
	Message string `json:"message"`
}
