package user

import "github.com/stretchr/testify/mock"

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) GetByID(id string) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockRepository) GetUserByEmail(email string) (*User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *mockRepository) Save(user User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockRepository) Login(user User) (*User, string, error) {
	args := m.Called(user)
	return args.Get(0).(*User), args.String(1), args.Error(2)
}
