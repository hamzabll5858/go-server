package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-server/pkg/controllers"
	"go-server/pkg/models"
	"go-server/pkg/middleware"
)

func InitRouter()  {
	models.InitDB()
	router := gin.Default()
	router.GET("/users", controllers.GetUsers)

	router.Use(middleware.VerifyCache())
	router.GET("/users/:id", controllers.GetUsersById)

	router.Run(viper.GetString("server.url"))
}