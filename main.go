package main

import (
	"crud/controllers"
	"crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.POST("/posts", controllers.PostCreate)
	r.Run()
}
