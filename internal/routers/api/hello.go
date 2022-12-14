package api

import "github.com/gin-gonic/gin"

type Hello struct{}

func NewHello() Hello {
	return Hello{}
}

func (h Hello) Index(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func (h Hello) Get(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"id": id,
	})
}

func (h Hello) Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    "Hello World",
	})
}

func (h Hello) PostForm(c *gin.Context) {
	data := c.PostForm("data")

	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h Hello) PostJson(c *gin.Context) {
	jsonData := make(gin.H)
	err := c.BindJSON(&jsonData)

	if err != nil {
		// error handle
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    jsonData["data"],
	})
}
