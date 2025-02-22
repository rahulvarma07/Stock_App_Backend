package models

type Response struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	TokenString string `json:"token,omitempty"`
}
