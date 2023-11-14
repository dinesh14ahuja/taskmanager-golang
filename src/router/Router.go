package router

import (
	"boilerplate-webserver-code/src/app/controller"
	"boilerplate-webserver-code/src/config"
	"boilerplate-webserver-code/src/middleware"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	router.Run(config.AppConfig.GetString("server.port"))
}

func NewRouter() *gin.Engine {
	router := gin.New()
	resource := router.Group("/task")
	resource.Use(middleware.LogRequestMiddleware())
	{
		resource.POST("", controller.NewTask)
		resource.GET("", controller.GetTasks)
		resource.DELETE("/:tasknumber", controller.DeleteTask)
		resource.PATCH("/:tasknumber", controller.EditTask)
	}
	return router
}
