package api

import (
	"gin-api/internal/pkg/app"
	"gin-api/internal/pkg/errors"
	"gin-api/internal/requests"

	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) Register(c *gin.Context) {
	params := requests.RegisterUserRequest{}
	response := app.NewResponse(c)
	valid, err := app.BindAndValidation(c, &params)

	if !valid {
		errorResponse := errors.InvalidParams.WithDetails(err...)
		response.MakeErrorResponse(errorResponse)

		return
	}

	response.MakeResponse(params)
}

func (u User) Login(c *gin.Context) {
	// ...
}
