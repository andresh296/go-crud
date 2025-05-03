package middleware

import (
	"net/http"
	"testing"

	schema "github.com/andresh296/go-crud/internal/platform/schema" 
	"github.com/test-go/testify/assert"
)

func TestJsonValidator(t *testing.T) {
    validators, err := schema.NewValidator(&schema.DefaultFileReader{})
    assert.NoError(t, err)
    builder := NewMiddlewareValidator(validators)

    tests := []struct {
        name       string
        inputJSON  string
        wantStatus int
        wantBody   string
    }{
        {
            name: "valid_login",
            inputJSON: `{
                "email": "test@example.com",
                "password": "password123"
            }`,
            wantStatus: 200,
            wantBody:   "",
        },
        {
            name: "invalid_email",
            inputJSON: `{
                "email": "not_an_email",
                "password": "password123"
            }`,
            wantStatus: 400,
            wantBody:   "validation failed",
        },
        {
            name: "missing_password",
            inputJSON: `{
                "email": "test@example.com"
            }`,
            wantStatus: 400,
            wantBody:   "validation failed",
        },
        {
            name: "invalid_json_format",
            inputJSON: `{
                "email": "test@example.com",,,
            }`,
            wantStatus: 400,
            wantBody:   "invalid JSON format",
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            
            c, w := createMockContext(http.MethodPost, "/login", []byte(test.inputJSON))

            
            builder.WithValidateLogin()(c)

    
            assert.Equal(t, test.wantStatus, w.Code, "Test case: "+test.name)

           
            if test.wantBody != "" {
                assert.Contains(t, w.Body.String(), test.wantBody, "Test case: "+test.name)
            }
        })
    }
}
