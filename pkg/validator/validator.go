package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validator defines the interface for struct validation.
type Validator interface {
	Validate(i interface{}) error
}

type customValidator struct {
	validator *validator.Validate
}

// NewValidator creates a new instance of the Validator.
func NewValidator() Validator {
	return &customValidator{
		validator: validator.New(),
	}
}

// Validate validates a struct using go-playground/validator tags.
func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
