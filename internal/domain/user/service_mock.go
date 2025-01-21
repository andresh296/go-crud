package user

import "github.com/stretchr/testify/mock"

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) GetByID(id string) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockRepository) GetUserByEmail(email string) (*User,error) {
	args := m.Called(email)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockRepository) Save(user User) error {
	args := m.Called(user)
	return args.Error(0)
}