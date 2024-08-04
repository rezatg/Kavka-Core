package user

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestIsValidUsername(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", usernameValidation)

	users := []UpdateProfileValidation{
		{Name: "reza", LastName: "None", Username: "rezatg"},
		{Name: "reza", LastName: "None", Username: "reza_tg"},
		{Name: "reza", LastName: "None", Username: "2rezatg"},
		{Name: "reza", LastName: "None", Username: ""},
		{Name: "reza", LastName: "None", Username: "rez"},
		{Name: "reza", LastName: "None", Username: "&^%$(*)"},
	}

	for _, user := range users {
		var err error = validate.Struct(user)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Validation passed for user:", user.Username)
		}
	}
}
