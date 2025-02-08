package api

import (
	domain "github.com/andresh296/go-crud/internal/domain/user"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

// Login implements user.Service.
func (m *MockService) Login(user domain.User) (*domain.User, error) {
	args := m.Called(&user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockService) GetByID(id string) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1) 
}

func (m *MockService) GetUserByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.User), args.Error(1) 
}


func (m *MockService) Save(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}
