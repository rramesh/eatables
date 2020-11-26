package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

//ValidationError holds the validator's Field Error
type ValidationError struct {
	validator.FieldError
}

//Error returns the formatted error message after validation
func (v *ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field Validation on '%s' failed on the tag '%s'",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

//ValidationErrors holds the slice of ValidationError
type ValidationErrors []ValidationError

//Errors returns the validation errors slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// Validation contains
type Validation struct {
	validate *validator.Validate
}

// NewValidation returns a new Validation
func NewValidation() *Validation {
	v := validator.New()
	v.RegisterValidation("uuid", validateUUID)
	return &Validation{v}
}

//Validate UUID
func validateUUID(fl validator.FieldLevel) bool {
	//UUID must be 12 length alphanumeric
	re := regexp.MustCompile(`[a-zA-Z0-9]{12}`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

// Validate validates a given structure
func (v *Validation) Validate(i interface{}) ValidationErrors {
	err := v.validate.Struct(i)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)
	if len(errs) == 0 {
		return nil
	}
	var returnErrs []ValidationError
	for _, err := range errs {
		//Cast FieldError into ValidationError
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}
	return returnErrs
}
