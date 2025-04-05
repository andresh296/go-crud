package middleware

import (
	schema "github.com/andresh296/go-crud/internal/platform/schema"
	"github.com/gin-gonic/gin"
	"github.com/kaptinlin/jsonschema"
)

type Builder struct {
	Validators *schema.Validators
}


func NewMiddlewareValidator(validators *schema.Validators) *Builder {
	
	return &Builder{
		Validators: validators,
	}
}

func (b *Builder) WithValidateLogin() gin.HandlerFunc {
	return b.jsonValidator(b.Validators.LoginValidator)
}

func (b *Builder) WithValidateRegister() gin.HandlerFunc {
	return b.jsonValidator(b.Validators.RegisterValidator)
}

func (b *Builder) jsonValidator(schema *jsonschema.Schema) gin.HandlerFunc {
	
	return func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"error": "invalid JSON format",
			})
			c.Abort()
			return
		}

		if err := schema.Validate(data); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
    
        
    }
}