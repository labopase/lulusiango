package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Validator struct {
	*validator.Validate
}

type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value,omitempty"`
	Message string `json:"message"`
}

func NewValidator() *Validator {
	v := validator.New()

	v.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		_, err := uuid.Parse(fl.Field().String())
		return err == nil
	})

	return &Validator{v}
}

func (v *Validator) ValidateStruct(s interface{}) []ValidationError {
	err := v.Struct(s)
	if err == nil {
		return nil
	}

	var errors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, ValidationError{
			Field:   getJSONFieldName(err),
			Tag:     err.Tag(),
			Value:   err.Param(),
			Message: getErrorMessage(err),
		})
	}
	return errors
}

func getJSONFieldName(err validator.FieldError) string {
	field := err.Field()
	return strings.ToLower(field[:1]) + field[1:]
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s", err.Field(), err.Param())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", err.Field(), err.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", err.Field(), err.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", err.Field(), err.Param())
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", err.Field())
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", err.Field(), err.Param())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", err.Field())
	case "alpha":
		return fmt.Sprintf("%s must contain only alphabetic characters", err.Field())
	case "alphanum":
		return fmt.Sprintf("%s must contain only alphanumeric characters", err.Field())
	case "numeric":
		return fmt.Sprintf("%s must be a valid number", err.Field())
	default:
		return fmt.Sprintf("%s failed on %s validation", err.Field(), err.Tag())
	}
}
