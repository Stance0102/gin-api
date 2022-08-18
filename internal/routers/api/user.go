package api

import (
	"github.com/Stance0102/gin-api/pkg/errors"

	"github.com/Stance0102/gin-api/internal/pkg/app"
	"github.com/Stance0102/gin-api/internal/requests"
	"github.com/Stance0102/gin-api/pkg/app"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) Register(c *gin.Context) {
	params := requests.RegisterUserRequest{}
	response := app.NewError(c)
	valid, err := app.BindAndValidation(c, &params)

	if !valid {
		errorResponse := errors.InvalidParams.WithDetails(err...)
		response.MakeErrorResponse(errorResponse)

		return
	}

	response.MakeErrorResponse(params)
}

func (u User) Login(c *gin.Context) {
	// ...
}
