package middleware

import (
	"net/http"

	"github.com/andresh296/go-crud/config"
	"github.com/andresh296/go-crud/internal/platform/security"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(cfg config.JWTConfig) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
            return
        }

        claims, err := security.ValidateJWT(tokenString, cfg.SecretKey)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            return
        }

        // Guardar claims en el contexto para uso posterior
        c.Set("user_id", claims.ID)
        c.Set("user_email", claims.Email)
        
        c.Next()
    }
}