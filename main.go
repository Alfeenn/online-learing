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
		baseRoute.POST("/register", controller.Register)
		baseRoute.GET("/login", controller.UserSignIn)
	}
	admin := baseRoute.Group("/admin")
	{
		admin.GET("/course", middleware.Authorize("course", "read", enforcer), controller.FindAll)
		admin.POST("/course", controller.Create)
		admin.GET("/:id", middleware.Authorize("course", "read", enforcer), controller.Find)
		admin.PUT("/:id", middleware.Authorize("course", "write", enforcer), controller.Update)
		admin.POST("/:id", middleware.Authorize("course", "delete", enforcer), controller.Delete)
		admin.GET("/acl", middleware.Authorize("course", "read", enforcer), controller.GetAccessList)
	}
	user := baseRoute.Group("/user", middleware.AuthJWT())
	{
		user.GET("", middleware.Authorize("course", "read", enforcer), controller.FindAll)
		user.GET("/:id", middleware.Authorize("course", "read", enforcer), controller.Find)
		user.GET("/class/:id", middleware.Authorize("class", "write", enforcer), controller.Find)
		user.GET("/acl", middleware.Authorize("course", "read", enforcer), controller.GetAccessList)
	}
	engine.Run("localhost:8000")
}
