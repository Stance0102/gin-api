package app

import (
	"net/http"

	"github.com/Stance0102/gin-api/pkg/errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

func (r *Response) MakeResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	r.Ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (r *Response) MakeErrorResponse(error *errors.Error) {
	newResponse := gin.H{
		"success": false,
		"message": error.GetMessage(),
		"details": error.GetDetails(),
	}

	r.Ctx.JSON(error.GetCode(), newResponse)
}
