package security

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "github.com/golang-jwt/jwt/v5"
)

// Test para verificar la generación exitosa de un token
func TestGenerateJWT_Success(t *testing.T) {
    // Arrange
    id := "123"
    email := "test@example.com"
    secretKey := "test-secret-key"
    expiration := time.Hour

    // Act
    token, err := GenerateJWT(id, email, secretKey, expiration)

    // Assert
    assert.NoError(t, err)
    assert.NotEmpty(t, token)
}

// Test para verificar el contenido del token generado
func TestGenerateJWT_Content(t *testing.T) {
    // Arrange
    id := "123"
    email := "test@example.com"
    secretKey := "test-secret-key"
    expiration := time.Hour

    // Act
    token, err := GenerateJWT(id, email, secretKey, expiration)
    assert.NoError(t, err)

    claims, err := ValidateJWT(token, secretKey)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, id, claims.ID)
    assert.Equal(t, email, claims.Email)
    assert.NotNil(t, claims.ExpiresAt)
}

// Test para verificar la validación de un token válido
func TestValidateJWT_ValidToken(t *testing.T) {
    // Arrange
    token, err := GenerateJWT("123", "test@example.com", "test-secret-key", time.Hour)
    assert.NoError(t, err)

    // Act
    claims, err := ValidateJWT(token, "test-secret-key")

    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, claims)
    assert.Equal(t, "123", claims.ID)
    assert.Equal(t, "test@example.com", claims.Email)
}

// Test para verificar el rechazo con clave secreta incorrecta
func TestValidateJWT_InvalidSecret(t *testing.T) {
    // Arrange
    token, err := GenerateJWT("123", "test@example.com", "correct-key", time.Hour)
    assert.NoError(t, err)

    // Act
    claims, err := ValidateJWT(token, "wrong-key")

    // Assert
    assert.Error(t, err)
    assert.Nil(t, claims)
}

// Test para verificar el rechazo de un token expirado
func TestValidateJWT_ExpiredToken(t *testing.T) {
    // Arrange
    token, err := GenerateJWT("123", "test@example.com", "test-secret-key", -time.Hour)
    assert.NoError(t, err)

    // Act
    claims, err := ValidateJWT(token, "test-secret-key")

    // Assert
    assert.Error(t, err)
    assert.Nil(t, claims)
    assert.ErrorIs(t, err, jwt.ErrTokenExpired)
}

// Test para verificar el rechazo de un token malformado
func TestValidateJWT_MalformedToken(t *testing.T) {
    // Act
    claims, err := ValidateJWT("este.token.no.es.valido", "test-secret-key")

    // Assert
    assert.Error(t, err)
    assert.Nil(t, claims)
}

// Test para verificar el rechazo de un token vacío
func TestValidateJWT_EmptyToken(t *testing.T) {
    // Act
    claims, err := ValidateJWT("", "test-secret-key")

    // Assert
    assert.Error(t, err)
    assert.Nil(t, claims)
}