package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	service := NewService(mockRepo)
	mockRepo.On("GetByID", "123").Return(expectedUser, nil)

	user, err := service.GetByID("123")
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}