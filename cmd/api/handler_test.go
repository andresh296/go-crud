package api

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	domain "github.com/andresh296/go-crud/internal/domain/user"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByEmail_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockservice := new(MockService)
	handler := New(mockservice)

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
	handler := New(mockservice)

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
	handler := New(mockservice)

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
	handler := New(mockservice)

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
	handler := New(mockservice)

	expecteduser := &domain.User{}

	mockservice.On("GetUserByEmail", "test@email.com").Return(expecteduser, domain.ErrDuplicateUser)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "email", Value: "test@email.com"}}

	handler.GetUserByEmail()(c)

	assert.Equal(t, http.StatusAlreadyReported, w.Code)
	//assert.Equal(t,expecteduser.Email,actualUser.Email)
}

func TestSave_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockService)

	handler := New(mockService)

	expectedUser := domain.User{

		Name:     "Emmanuel",
		Email:    "emmago@gmail.com",
		Age:      27,
		Password: "123456789",
	}

	mockService.On("Save", expectedUser).Return(expectedUser, nil)

	body := `{"name": "Emmanuel", "age": 27, "email": "emmago@gmail.com", "password": "123456789"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	handler.Save()(c)
	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestSave_error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Crear un mock del servicio
	mockService := new(MockService)

	// Crear el handler con el mock
	handler := New(mockService)

	// Datos esperados de usuario después de guardarse
	expectedUser := domain.User{

		Name:     "Emmanuel",
		Email:    "emmago@gmail.com",
		Age:      27,
		Password: "123456789",
	}

	mockService.On("Save", expectedUser).Return(domain.User{}, domain.ErrUserCannotSave)

	body := `{"name": "Emmanuel", "age": 27, "email": "emmago@gmail.com", "password": "123456789"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	handler.Save()(c)

	assert.Equal(t, http.StatusFailedDependency, w.Code)

}

func TestSave_errorValidate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockService)
	handler := New(mockService)

	expectedUser := domain.User{

		Name:     "Emmanuel",
		Email:    "testValidate.com",
		Age:      27,
		Password: "123456789",
	}

	mockService.On("Save", expectedUser).Return(expectedUser, domain.ErrValidationUser)

	body := `{"name": "Emmanuel", "age": 27, "email": "testValidate.com", "password": "123456789"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	handler.Save()(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

}

func TestSave_errorJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockService)
	handler := New(mockService)

	expectedUser := domain.User{
		Name:     "Emmanuel",
		Email:    "emmago@gmail.com",
		Password: "123456789",
	}

	mockService.On("Save", expectedUser).Return(domain.User{}, ErrUnmarshalBody)

	body := `{"name": "Emmanuel", edad:20, "email": "emmago@gmail.com", "password": "123456789"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	handler.Save()(c)
	assert.Equal(t, http.StatusBadRequest, w.Code)

}
