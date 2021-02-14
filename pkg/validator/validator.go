package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

func (v Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}
