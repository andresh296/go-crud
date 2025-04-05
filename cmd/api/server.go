package api

import (
	domain "github.com/andresh296/go-crud/internal/domain/user"
	mid "github.com/andresh296/go-crud/internal/middleware"
	schema "github.com/andresh296/go-crud/internal/platform/schema"
	"github.com/gin-gonic/gin"
)


func routing(app *gin.Engine, dependencies *Dependencies) {
	userService := domain.NewService(dependencies.userRepo)
	handler := New(userService)

	validators, err := schema.NewValidator()
	if err != nil {
		panic(err)
	}
	validator := mid.NewMiddlewareValidator(validators)


    public := app.Group("/v1")
    {
        public.POST("/user/login", validator.WithValidateLogin(), handler.Login())
        public.POST("/user", validator.WithValidateRegister(), handler.Save())
    }

	protected := app.Group("/v1")
	protected.Use(mid.JWTAuthMiddleware(dependencies.config.JWT))
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
