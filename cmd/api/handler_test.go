package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andresh296/go-crud/config"
	domain "github.com/andresh296/go-crud/internal/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestGetUserByEmail_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockservice := new(MockService)
	handler := New(mockservice, &config.Config{})

	expecteduser := &domain.User{}

	mockservice.On("GetUserByEmail", "test@email.com").Return(expecteduser, domain.ErrGettingUserByEmail)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "email", Value: "test@email.com"}}

	handler.GetUserByEmail()(c)

	assert.Equal(t, http.StatusFailedDependency, w.Code)
}

func TestGetUserByEmail_Succes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockservice := new(MockService)
	handler := New(mockservice, &config.Config{})

	expecteduser := &domain.User{
		ID:    "1",
		Name:  "testemail",
		Email: "test@email.com",
	}

	mockservice.On("GetUserByEmail", "test@email.com").Return(expecteduser, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "email", Value: "test@email.com"}}

	handler.GetUserByEmail()(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserByID_Succes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockservice := new(MockService)
	handler := New(mockservice, &config.Config{})

	expecteduser := &domain.User{
		ID:    "1238",
		Name:  "testemail",
		Email: "test@email.com",
	}

	mockservice.On("GetByID", "1238").Return(expecteduser, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1238"}}

	handler.GetByID()(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserByID_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockservice := new(MockService)
	handler := New(mockservice, &config.Config{})

	mockservice.On("GetByID", "1238").Return(&domain.User{}, domain.ErrUserCannotGet)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1238"}}

	handler.GetByID()(c)

	assert.Equal(t, http.StatusFailedDependency, w.Code)
}

func TestGetUserByEmail_ErrorDuplicate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockservice := new(MockService)
	handler := New(mockservice, &config.Config{})

	expecteduser := &domain.User{}

	mockservice.On("GetUserByEmail", "test@email.com").Return(expecteduser, domain.ErrDuplicateUser)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "email", Value: "test@email.com"}}

	handler.GetUserByEmail()(c)
	assert.Equal(t, http.StatusAlreadyReported, w.Code)

}

func TestLogin_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockservice := new(MockService)
	cfg := &config.Config{
		JWT: config.JWTConfig{
			SecretKey:      "test-secret",
			ExpirationTime: 3600,
		},
	}
	handler := New(mockservice, cfg)

	expecteduser := &domain.User{
		Email:    "test@email.com",
		Password: "testpassword",
	}

	mockservice.On("Login", expecteduser).Return("test@email.com", nil)

	jsonBody := []byte(`{"email": "test@email.com", "password": "testpassword"}`)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, err := http.NewRequest("POST", "/v1/user/login", io.NopCloser(bytes.NewBuffer(jsonBody)))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	c.Request = req

	handler.Login()(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal("Error al parsear la respuesta:", err)
	}
	assert.Equal(t, expecteduser.Email, response["email"])
}
