package httpserver

import "gopkg.in/go-playground/validator.v9"

// Validator is a main validator
type Validator struct {
	validator *validator.Validate
}

// Validate validates struct
func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
