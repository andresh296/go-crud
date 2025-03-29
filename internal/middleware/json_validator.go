package middleware

import (
	"github.com/andresh296/go-crud/internal/platform/schema"
	"github.com/gin-gonic/gin"
	"github.com/kaptinlin/jsonschema"
)

type BuilderMiddleware struct {
	Validators *schema.Validators
}

func NewMiddlewareValidator(validators *schema.Validators) *BuilderMiddleware {
	return &BuilderMiddleware{
		Validators: validators,
	}
}

func (b *BuilderMiddleware) WithValidateLogin() gin.HandlerFunc {
	return b.jsonValidator(b.Validators.LoginValidator)
}

func (b *BuilderMiddleware) WithValidateRegister() gin.HandlerFunc {
	return b.jsonValidator(b.Validators.RegisterValidator)
}

func (b *BuilderMiddleware) jsonValidator(schema *jsonschema.Schema) gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}