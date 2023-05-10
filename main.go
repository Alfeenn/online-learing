package main

import (
	"github.com/Alfeenn/online-learning/app"
	"github.com/Alfeenn/online-learning/cmd"
	"github.com/Alfeenn/online-learning/controller"
	"github.com/Alfeenn/online-learning/middleware"
	"github.com/Alfeenn/online-learning/repository"
	"github.com/Alfeenn/online-learning/service"
	"github.com/gin-gonic/gin"
)

func main() {
	migrate, enforcer := cmd.MigrateCmd()
	if migrate {
		return
	}
	engine := gin.New()
	db := app.NewDB()
	repo := repository.NewRepository()
	service := service.NewService(repo, db)
	controller := controller.NewController(service)
	auth := middleware.NewMiddleware()
	baseRoute := engine.Group("/api", auth)
	{
		baseRoute.POST("/register", controller.Create)
		baseRoute.GET("/login", controller.UserSignIn)
	}
	user := baseRoute.Group("/user", middleware.AuthJWT())
	{
		user.GET("", middleware.Authorize("data", "read", enforcer), controller.FindAll)
		user.GET("/:id", middleware.Authorize("data", "read", enforcer), controller.Find)
		user.PUT("/:id", middleware.Authorize("data", "write", enforcer), controller.Update)
		user.POST("/:id", middleware.Authorize("data", "write", enforcer), controller.Delete)
		user.GET("/acl", middleware.Authorize("data", "read", enforcer), controller.GetAccessList)
	}
	engine.Run("localhost:8000")
}
