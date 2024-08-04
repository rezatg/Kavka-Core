package user

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// UpdateProfileValidation defines the structure for profile validation
type UpdateProfileValidation struct {
	Name     string `validate:"required,min=3,max=40"`
	LastName string `validate:"required,min=3,max=40"`
	Username string `validate:"required,username,min=4,max=20"`
}

// Compile the regex for username validation
var usernameRegex *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{4,20}$`)

// Custom validation function for username
func usernameValidation(fl validator.FieldLevel) bool {
	var _uname string = fl.Field().String()
	if _uname == "" {
		return true
	}

	return usernameRegex.MatchString(_uname)
}

// RegisterCustomValidations registers the custom validation functions
// func registerCustomValidations(validate *validator.Validate) {
// 	validate.RegisterValidation("username", usernameValidation)
// }
