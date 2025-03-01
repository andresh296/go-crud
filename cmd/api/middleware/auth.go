package middleware

import (
    "fmt"
    "net/http"

    "github.com/andresh296/go-crud/config"
    "github.com/andresh296/go-crud/internal/platform/security"
    "github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(cfg config.JWTConfig) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        fmt.Printf("Token recibido: %s\n", token)

        if len(token) > 7 && token[:7] == "Bearer " {
            token = token[7:]
        }

        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
            return
        }

        claims, err := security.ValidateJWT(token, cfg.SecretKey)
        if err != nil {
            fmt.Printf("Error validando token: %v\n", err)
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            return
        }

        c.Set("user_id", claims.ID)
        c.Set("user_email", claims.Email)
        c.Next()
    }
}