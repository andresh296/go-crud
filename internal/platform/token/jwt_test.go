package security

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "github.com/golang-jwt/jwt/v5"
)


func TestGenerateJWT_Success(t *testing.T) {

    id := "123"
    email := "test@example.com"
    secretKey := "test-secret-key"
    expiration := time.Hour


    token, err := GenerateJWT(id, email, secretKey, expiration)


    assert.NoError(t, err)
    assert.NotEmpty(t, token)
}


func TestGenerateJWT_Content(t *testing.T) {
  
    id := "123"
    email := "test@example.com"
    secretKey := "test-secret-key"
    expiration := time.Hour

 
    token, err := GenerateJWT(id, email, secretKey, expiration)
    assert.NoError(t, err)

    claims, err := ValidateJWT(token, secretKey)
    

    assert.NoError(t, err)
    assert.Equal(t, id, claims.ID)
    assert.Equal(t, email, claims.Email)
    assert.NotNil(t, claims.ExpiresAt)
}


func TestValidateJWT_ValidToken(t *testing.T) {

    token, err := GenerateJWT("123", "test@example.com", "test-secret-key", time.Hour)
    assert.NoError(t, err)

   
    claims, err := ValidateJWT(token, "test-secret-key")


    assert.NoError(t, err)
    assert.NotNil(t, claims)
    assert.Equal(t, "123", claims.ID)
    assert.Equal(t, "test@example.com", claims.Email)
}

func TestValidateJWT_InvalidSecret(t *testing.T) {
    
    token, err := GenerateJWT("123", "test@example.com", "correct-key", time.Hour)
    assert.NoError(t, err)

    
    claims, err := ValidateJWT(token, "wrong-key")

    
    assert.Error(t, err)
    assert.Nil(t, claims)
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
 
    token, err := GenerateJWT("123", "test@example.com", "test-secret-key", -time.Hour)
    assert.NoError(t, err)

   
    claims, err := ValidateJWT(token, "test-secret-key")

 
    assert.Error(t, err)
    assert.Nil(t, claims)
    assert.ErrorIs(t, err, jwt.ErrTokenExpired)
}


func TestValidateJWT_MalformedToken(t *testing.T) {
    
    claims, err := ValidateJWT("este.token.no.es.valido", "test-secret-key")

    assert.Error(t, err)
    assert.Nil(t, claims)
}

func TestValidateJWT_EmptyToken(t *testing.T) {
    
    claims, err := ValidateJWT("", "test-secret-key")


    assert.Error(t, err)
    assert.Nil(t, claims)
}