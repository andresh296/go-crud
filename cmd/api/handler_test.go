package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	domain "github.com/andresh296/go-crud/internal/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestGetUserByEmail_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockservice := new(MockService)
	handler := New(mockservice)

	expecteduser := &domain.User{}

	// Configurar el mock para devolver un error ya definido en `domain`
	mockservice.On("GetUserByEmail", "test@email.com").Return(expecteduser, domain.ErrGettingUserByEmail)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "email", Value: "test@email.com"}}

	handler.GetUserByEmail()(c)

	// Validar que devuelve un 424 en vez de 500
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

	// Configurar el mock para devolver un error ya definido en `domain`
	mockservice.On("GetUserByEmail", "test@email.com").Return(expecteduser, domain.ErrDuplicateUser)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "email", Value: "test@email.com"}}

	handler.GetUserByEmail()(c)

	//var actualUser domain.User

	assert.Equal(t, http.StatusAlreadyReported, w.Code)
	//assert.Equal(t,expecteduser.Email,actualUser.Email)
}
