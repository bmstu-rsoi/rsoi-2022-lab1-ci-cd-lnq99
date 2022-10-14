package model

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

type ValidationErrorResponse struct {
	Errors  ValidationErrorResponse_Errors `json:"errors,omitempty"`
	Message string                         `json:"message,omitempty"`
}

type ValidationErrorResponse_Errors struct {
	AdditionalProperties string `json:"-"`
	// AdditionalProperties map[string]string `json:"-"`
}
