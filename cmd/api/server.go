package api

import (
	domain "github.com/andresh296/go-crud/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func routing(app *gin.Engine, dependencies *Dependencies) {
	userService := domain.NewService(dependencies.user)
	handler := New(userService)

	app.GET("/v1/user", handler.GetAll())
	app.POST("/v1/user", handler.Save())
}

func Boostrap(app *gin.Engine) {
	dependencies := initDependencies()
	routing(app, dependencies)
}