package main

import (
	"feed/controllers"
	middleware "feed/middlewares"
	"feed/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:username", middleware.FreeFeedMiddleware(), controllers.FindUser)
		v1.PATCH("/users/:username", middleware.FreeFeedMiddleware(), controllers.UpdateBirthday)
	}

	r.Run()
}
