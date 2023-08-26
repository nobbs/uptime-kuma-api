package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

// Validator returns a validator instance. This is a singleton, as the validator provides caching
// capabilities built-in.
func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}

	return validate
}

// ValidateStruct validates a struct and returns an error if the validation fails.
func ValidateStruct(s any) error {
	return GetValidator().Struct(s)
}

// ValidateVar validates a variable and returns an error if the validation fails.
func ValidateVar(s any, tag string) error {
	return GetValidator().Var(s, tag)
}
