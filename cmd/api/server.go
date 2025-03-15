package api

import (
	"github.com/andresh296/go-crud/internal/middleware"

	domain "github.com/andresh296/go-crud/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func routing(app *gin.Engine, dependencies *Dependencies) {
	userService := domain.NewService(dependencies.userRepo)
	handler := New(userService)

	app.POST("/v1/user", handler.Save())
	app.POST("/v1/user/login", handler.Login())

	protected := app.Group("/v1")
	protected.Use(middleware.JWTAuthMiddleware(dependencies.config.JWT))
	{
		protected.GET("/user/id/:id", handler.GetByID())
		protected.GET("/user/email/:email", handler.GetUserByEmail())
	}
}

func Boostrap(app *gin.Engine) {
	dependencies := initDependencies()
	if dependencies == nil {
		panic("dependencies not initialized")
	}
	routing(app, dependencies)
}
