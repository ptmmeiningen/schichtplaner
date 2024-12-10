package models

// APIResponse Standard API Antwortformat
type APIResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message,omitempty" example:"Operation erfolgreich"`
	Error   string      `json:"error,omitempty" example:"Fehlermeldung"`
	Data    interface{} `json:"data,omitempty"`
}
