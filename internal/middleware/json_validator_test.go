package middleware

import (
	"testing"

	schema "github.com/andresh296/go-crud/internal/platform/schema"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJsonValidator(t *testing.T) {
    validators, err := schema.NewValidator()
    assert.NoError(t, err)
    
    builder := NewMiddlewareValidator(validators)

    tests := []struct {
        name       string
        inputJSON  map[string]interface{}
        wantStatus int
    }{
        {
            name: "valid_login",
            inputJSON: map[string]interface{}{
                "email": "test@example.com",
                "password": "password123",
            },
            wantStatus: 200,
        },
        {
            name: "invalid_email",
            inputJSON: map[string]interface{}{
                "email": "not_an_email",
                "password": "password123",
            },
            wantStatus: 400,
        },
        {
            name: "missing_password",
            inputJSON: map[string]interface{}{
                "email": "test@example.com",
            },
            wantStatus: 400,
        },
    }

    for _, test := range tests {
        c := &MockGinContext{
            mockJSON: test.inputJSON,
            mockData: make(map[string]interface{}),
        }

   
        builder.WithValidateLogin()(c.ToGinContext().(*gin.Context))
	

        assert.Equal(t, test.wantStatus, c.mockStatus, "Test case: "+test.name)
    }
}
