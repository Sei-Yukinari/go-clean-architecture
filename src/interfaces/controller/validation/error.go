package validation

type ModelValidationError struct {
	Message string
}

func NewModelValidationError(message string) *ModelValidationError {
	return &ModelValidationError{Message: message}
}

func (e *ModelValidationError) Error() string {
	return e.Message
}
