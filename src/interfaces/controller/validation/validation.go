package validation

import (
	"fmt"
	"go-clean-architecture/src/infrastructure/logger"

	validatorV10 "github.com/go-playground/validator/v10"
)

var validatorInstance *validatorV10.Validate

const validationErrorPrefix = "Validation apperror: "

type validator struct{}

type Validator interface {
	Validate(interface{}) error
}

func NewValidator() Validator {
	return validator{}
}

func init() {
	validatorInstance = validatorV10.New()
}

func (v validator) Validate(value interface{}) error {
	err := validatorInstance.Struct(value)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validatorV10.InvalidValidationError); ok {
		logger.Warn(err)
		return nil
	}

	errStr := validationErrorPrefix
	for _, err := range err.(validatorV10.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errStr += fmt.Sprintf("field '%v' is required, ", err.Field())
		case "min":
			errStr += fmt.Sprintf("field '%v' is too short, ", err.Field())
		case "max":
			errStr += fmt.Sprintf("field '%v' is too long, ", err.Field())
		case "gte":
		case "lte":
			errStr += fmt.Sprintf("field '%v' is invalid, ", err.Field())
		case "alphanum":
			errStr += fmt.Sprintf("field '%v' is alpha or number ", err.Field())
		case "email":
			errStr += fmt.Sprintf("field '%v' is invalid email format ", err.Field())
		default:
			errStr += fmt.Sprintf("field '%v' validation failed because tag '%v', ", err.Field(), err.Tag())
		}
	}

	return NewModelValidationError(errStr)
}
