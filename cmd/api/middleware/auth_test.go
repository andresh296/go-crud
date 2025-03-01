package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/andresh296/go-crud/config"
	"github.com/andresh296/go-crud/internal/platform/security"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTest() (*gin.Engine, config.JWTConfig) {
	cfg := config.JWTConfig{
		SecretKey: "test-secret-key",
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(JWTAuthMiddleware(cfg))

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	return r, cfg
}

func TestJWTAuthMiddleware_MissingToken(t *testing.T) {
	r, _ := setupTest()

	// Crear request sin token
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()


	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestJWTAuthMiddleware_VerifyClaimsInContext(t *testing.T) {
    
    cfg := config.JWTConfig{
        SecretKey: "test-secret-key",
    }

    // Crear token con claims específicos
    id := "123"
    email := "test@example.com"

    
    token, err := security.GenerateJWT(id, email, cfg.SecretKey, 1*time.Hour)
    if err != nil {
        t.Fatalf("Error generando token: %v", err)
    }
    
    
    var capturedUserID string
    var capturedEmail string


    gin.SetMode(gin.TestMode)
    r := gin.New()
    r.Use(JWTAuthMiddleware(cfg))

    r.GET("/test", func(c *gin.Context) {
    
        userID, exists := c.Get("user_id")
        if !exists {
            t.Error("La clave 'user_id' no está presente en el contexto")
            c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found"})
            return
        }
        if id, ok := userID.(string); ok {
            capturedUserID = id
        } else {
            t.Error("La clave 'user_id' no es una cadena")
            c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id wrong type"})
            return
        }

    
        userEmail, exists := c.Get("user_email") 
        if !exists {
            t.Error("La clave 'user_email' no está presente en el contexto")
            c.JSON(http.StatusInternalServerError, gin.H{"error": "user_email not found"})
            return
        }
        if email, ok := userEmail.(string); ok {
            capturedEmail = email
        } else {
            t.Error("La clave 'user_email' no es una cadena")
            c.JSON(http.StatusInternalServerError, gin.H{"error": "user_email wrong type"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"status": "success"})
    })


    req := httptest.NewRequest(http.MethodGet, "/test", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    
    if w.Code != http.StatusOK {
        fmt.Printf("Response Body: %s\n", w.Body.String())
    }

    
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, id, capturedUserID, "El ID capturado no coincide con el original")
    assert.Equal(t, email, capturedEmail, "El Email capturado no coincide con el original")

    
    claims, err := security.ValidateJWT(token, cfg.SecretKey)
    assert.NoError(t, err, "Error validando el token")
    assert.NotNil(t, claims, "Los claims no deberían ser nil")

    assert.NotEmpty(t, capturedUserID, "El ID capturado no debería estar vacío")
    assert.NotEmpty(t, capturedEmail, "El Email capturado no debería estar vacío")
}

func TestJWTAuthMiddleware_NoToken(t *testing.T) {
	r, _ := setupTest()

	
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	
	r.ServeHTTP(w, req)


	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Token no proporcionado", response["error"])
}

func TestJWTAuthMiddleware_InvalidToken(t *testing.T) {
	r, _ := setupTest()


	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "invalid.token.here")
	w := httptest.NewRecorder()

	
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Token inválido", response["error"])
}

func TestJWTAuthMiddleware_ValidToken(t *testing.T) {
	r, cfg := setupTest()



	token, err := security.GenerateJWT("test-id-jwt", "test-email-jwt", cfg.SecretKey, time.Hour)
	assert.NoError(t, err)
	fmt.Printf("Token generado: %s\n", token)

	
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token) 
	w := httptest.NewRecorder()

	
	r.ServeHTTP(w, req)

	
	if w.Code != http.StatusOK {
		fmt.Printf("Response Body: %s\n", w.Body.String())
	}

	
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err = json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["status"])
}
