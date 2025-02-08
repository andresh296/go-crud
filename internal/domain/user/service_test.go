package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByID_Success(t *testing.T) {
	mockRepo := &mockRepository{}
	expectedUser := &User{
		ID: "123",
		Name: "test",
		Age: 20,
		Email: "email@test",
		Password: "12345",
	}

	mockRepo.On("GetByID", "123").Return(expectedUser, nil)

	service := NewService(mockRepo)

	user, err := service.GetByID("123")
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestGetByID_ErrorNotFound(t *testing.T) {
	mockRepo := &mockRepository{}

	mockRepo.On("GetByID", "123").Return(&User{}, ErrUserCannotGet)

	service := NewService(mockRepo)

	_, err := service.GetByID("123")
	assert.Equal(t, ErrUserCannotGet, err)
}

func TestGetUserByEmail_Success(t *testing.T) {
	mockRepo := &mockRepository{}
	expectedUser := &User{
		ID: "123",
		Name: "test",
		Age: 20,
		Email: "email@test",
		Password: "12345",
	}

	mockRepo.On("GetUserByEmail", "email@test").Return(expectedUser, nil)

	service := NewService(mockRepo)
	user, err := service.GetUserByEmail("email@test")
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestGetUserByEmail_ErrorNotFound(t *testing.T) {
	mockRepo := &mockRepository{}

	mockRepo.On("GetUserByEmail", "email@test").Return(&User{}, ErrNotFoundUserByEmail)

	service := NewService(mockRepo)
	_, err := service.GetUserByEmail("email@test")

	assert.Equal(t, ErrNotFoundUserByEmail, err)
}

func TestSave_Success(t *testing.T) {
	mockRepo := &mockRepository{}
	expectedUser := User{
		ID: "123",
		Name: "test",
		Age: 20,
		Email: "email@test",
		Password: "12345",
	}
	expectedUser.hashPassword()

	userMatched := mock.MatchedBy(func (actual User) bool {
		return compareUser(expectedUser, actual)
	})

	mockRepo.On("Save", userMatched).Return(nil)

	service := NewService(mockRepo)
	user, err := service.Save(expectedUser)

	assert.Nil(t, err)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Age, user.Age)
	assert.Equal(t, expectedUser.Email, user.Email)
}

func TestSave_ErrorCannotSaveUser(t *testing.T) {
	mockRepo := &mockRepository{}
	expectedUser := User{
		ID: "123",
		Name: "test",
		Age: 20,
		Email: "email@test",
		Password: "12345",
	}
	expectedUser.hashPassword()

	userMatched := mock.MatchedBy(func (actual User) bool {
		return compareUser(expectedUser, actual)
	})

	mockRepo.On("Save", userMatched).Return(ErrUserCannotSave)

	service := NewService(mockRepo)
	_, err := service.Save(expectedUser)

	assert.Equal(t, ErrUserCannotSave, err)
}

func compareUser(expected, actual User) bool {
	return expected.Name == actual.Name && expected.Age == actual.Age && expected.Email == actual.Email
}
