package middleware

import (
	"encoding/json"

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
	return b.jsonValidator(b.Validators.RegisterValidator )
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

    
        result := schema.Validate(data)
        if !result.IsValid() {
            details, err := json.MarshalIndent(result.ToList(), "", "  ")
            if err != nil {
                c.JSON(500, gin.H{
                    "error": "internal server error",
                })
                c.Abort()
                return
            }
            c.JSON(400, gin.H{
                "error": "validation failed",
                "details": string(details),
            })
            c.Abort()
            return
        }

        c.Set("validatedData", data)
        c.Next()
    }
}