package user

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type mySuite struct {
	suite.Suite

	v *validator.Validate
}

func (m *mySuite) SetupTest() {
	m.v = validator.New()

	var err error = registerCustomValidations(m.v)
	assert.NoError(m.T(), err)
}

// Validness
func (m *mySuite) TestUsernameValidity() {
	testCases := []struct {
		Name     string `validate:"required,min=3,max=40"`
		LastName string `validate:"required,min=3,max=40"`
		Username string `validate:"required,username,min=4,max=20"`

		ExpectedError string
	}{
		{
			Name:     "reza",
			LastName: "None",
			Username: "rezatg",
		},
		{
			Name:     "reza",
			LastName: "None",
			Username: "reza_tg",
		},
		{
			Name:     "reza",
			LastName: "None",
			Username: "2rezatg",
		},
		{
			Name:     "reza",
			LastName: "None",
			Username: "",
		},
		{
			Name:     "reza",
			LastName: "None",
			Username: "rez",
		},
		{
			Name:     "reza",
			LastName: "None",
			Username: "&^%$(*)",
		},
	}

	for _, user := range testCases {
		var err error = m.v.Struct(user)

		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				assert.EqualError(m.T(), err, fmt.Sprintf("username is not valid: %s", user.Username))
			}
		} else {
			assert.NoError(m.T(), err)
		}
	}
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(mySuite))
}
