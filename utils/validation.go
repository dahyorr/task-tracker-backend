package utils

import "github.com/go-playground/validator"

var validate = validator.New()

type ValidationError struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func ValidateData(data interface{}) []ValidationError {
	validationErrors := []ValidationError{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ValidationError

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
