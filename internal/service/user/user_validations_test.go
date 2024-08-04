package user

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// MyUserValidationTestSuite is a test suite for validating user information.
type MyUserValidationTestSuite struct {
	suite.Suite

	v *validator.Validate
}

// SetupTest sets up the test environment and initializes the validator.
func (s *MyUserValidationTestSuite) SetupTest() {
	s.v = validator.New()

	var err error = registerCustomValidations(s.v)
	assert.NoError(s.T(), err)
}

// Validness tests the validity of user information including name, last name, and username.
func (s *MyUserValidationTestSuite) Validness() {
	testCases := []struct {
		Name     string `validate:"required,min=3,max=40"`
		LastName string `validate:"required,min=3,max=40"`
		Username string `validate:"required,username,min=4,max=20"`

		Valid bool
	}{
		// Test cases with different user information
		{
			Name:     "john",
			LastName: "None",
			Username: "john",
		},
		{
			Name:     "john",
			LastName: "None",
			Username: "john_doe",
		},
		{
			Name:     "john",
			LastName: "None",
			Username: "2john",
		},
		{
			Name:     "john",
			LastName: "None",
			Username: "",
		},
		{
			Name:     "john",
			LastName: "None",
			Username: "joh",
		},
		{
			Name:     "john",
			LastName: "None",
			Username: "&^%$(*)",
		},
	}

	// Loop through test cases and validate user information
	for _, user := range testCases {
		var err error = s.v.Struct(user)

		if err != nil {
			// Check for validation errors and display appropriate error messages
			for _, err := range err.(validator.ValidationErrors) {
				assert.EqualError(s.T(), err, fmt.Sprintf("username is not valid: %s", user.Username))
				// require.Equal(s.T(), (err != nil) == user.Valid, user.Username)
			}
		} else {
			assert.NoError(s.T(), err)
		}
	}
}

// TestRunSuite runs the test suite for user validation.
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(MyUserValidationTestSuite))
}
