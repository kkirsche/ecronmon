package models

// Health represents the current health of the application
type Health struct {
	Healthy bool   `json:"healthy"`
	Message string `json:"message"`
}
