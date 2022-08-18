package routers

import (
	"github.com/Stance0102/gin-api/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	hello := api.NewHello()
	user := api.NewUser()

	router.GET("/", hello.Index)
	router.GET("/test", hello.Test)
	router.POST("/form", hello.PostForm)
	router.POST("/json", hello.PostJson)
	router.GET("/:id", hello.Get)

	router.POST("/register", user.Register)
	router.POST("/login", user.Login)

	return router
}
