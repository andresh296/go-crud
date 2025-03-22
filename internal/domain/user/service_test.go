package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestGetByID_Success(t *testing.T) {
	mockRepo := &mockRepository{}
	expectedUser := &User{
		ID:       "123",
		Name:     "test",
		Age:      20,
		Email:    "email@test",
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
		ID:       "123",
		Name:     "test",
		Age:      20,
		Email:    "email@test",
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
		ID:       "123",
		Name:     "test",
		Age:      20,
		Email:    "email@test",
		Password: "12345",
	}
	expectedUser.hashPassword()

	userMatched := mock.MatchedBy(func(actual User) bool {
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
		ID:       "123",
		Name:     "test",
		Age:      20,
		Email:    "email@test",
		Password: "12345",
	}
	expectedUser.hashPassword()

	userMatched := mock.MatchedBy(func(actual User) bool {
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

func compareUserLogin(expected, actual User) bool {
	err := bcrypt.CompareHashAndPassword([]byte(actual.Password), []byte(expected.Password))
	if err != nil {
		return false
	}
	return expected.Email == actual.Email
}

func TestLogin_Success(t *testing.T) {
	mockRepo := &mockRepository{}
	expectedUser := User{
		Email:    "email@test",
		Password: "12345",
	}
	expectedUser.hashPassword()

	userMatched := mock.MatchedBy(func(actual User) bool {
		return compareUserLogin(expectedUser, actual)
	})

	mockRepo.On("GetUserByEmail", "email@test").Return(&expectedUser, nil)
	mockRepo.On("Login", userMatched).Return(&expectedUser, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaGFuIjoiMjMwfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", nil)

	service := NewService(mockRepo)
	user, token, err := service.Login(User{
		Email:    "email@test",
		Password: "12345",
	})

	assert.Nil(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Age, user.Age)
	assert.Equal(t, expectedUser.Email, user.Email)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := &mockRepository{}

	mockRepo.On("GetUserByEmail", "email@test").Return(nil, ErrNotFoundUserByEmail)

	service := NewService(mockRepo)
	user, token, err := service.Login(User{
		Email:    "email@test",
		Password: "12345",
	})

	assert.Equal(t, ErrNotFoundUserByEmail, err)
	assert.Nil(t, user)
	assert.Empty(t, token)
}

func TestLogin_PasswordNotMatch(t *testing.T) {
	mockRepo := &mockRepository{}
	expectedUser := User{
		Email:    "email@test",
		Password: "12345",
	}
	expectedUser.hashPassword()

	mockRepo.On("GetUserByEmail", "email@test").Return(&expectedUser, ErrUserCannotLogin)

	service := NewService(mockRepo)
	user, token, err := service.Login(User{
		Email:    "email@test",
		Password: "123456",
	})

	assert.Equal(t, ErrUserCannotLogin, err)
	assert.Nil(t, user)
	assert.Empty(t, token)
}

func TestLogin_ErrorToken(t *testing.T) {
    mockRepo := &mockRepository{}
    expectedUser := User{
        Email:    "email@test",
        Password: "12345",
    }
    expectedUser.hashPassword()

    userMatched := mock.MatchedBy(func(actual User) bool {
        return compareUserLogin(expectedUser, actual)
    })

    mockRepo.On("GetUserByEmail", "email@test").Return(&expectedUser, nil)
    mockRepo.On("Login", userMatched).Return(expectedUser, "", ErrUserCannotLogin)

    service := NewService(mockRepo)
    _, token, err := service.Login(User{
        Email:    "email@test",
        Password: "12345",
    })
	
	assert.Nil(t, err)
	assert.NotEqual(t,token , "")
}

