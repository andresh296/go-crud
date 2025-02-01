package api

import (
	domain "github.com/andresh296/go-crud/internal/domain/user"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetByID(id string) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1) // Devuelve el valor simulado y el error
}

// GetUserByEmail simula la obtención de un usuario por email
func (m *MockService) GetUserByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.User), args.Error(1) // Devuelve el valor simulado y el error
}

// Save simula la creación o actualización de un usuario
func (m *MockService) Save(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1) //// Devuelve el valor simulado y el error
}
