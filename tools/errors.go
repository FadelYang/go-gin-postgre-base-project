package tools

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Errors []FieldError `json:"errors"`
}

func NewValidationError(field, msg string) *ValidationError {
	return &ValidationError{
		Errors: []FieldError{
			{Field: field, Message: msg},
		},
	}
}

func (v *ValidationError) Error() string {
	return "validation error"
}
