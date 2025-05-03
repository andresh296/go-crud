package middleware

import (
	//"errors"
	"errors"
	"testing"

	"github.com/test-go/testify/assert"
)

func TestNewValidator_MockConfiguration(t *testing.T) {
    mockReader := new(FileReaderMock)
    mockReader.On("ReadJsonSchema", "login_schema.json").Return([]byte(`{}`), nil)
    mockReader.On("ReadJsonSchema", "register_schema.json").Return([]byte(`{}`), nil)

    _, err := NewValidator(mockReader)
   
    assert.Nil(t, err)
    mockReader.AssertExpectations(t)
}

func TestNewValidator(t *testing.T) {
	mockReader := new(FileReaderMock)
	mockReader.On("ReadJsonSchema", "login_schema.json").Return([]byte(`{}`), nil)
	mockReader.On("ReadJsonSchema", "register_schema.json").Return([]byte(`{}`), nil)

    validators, err := NewValidator(mockReader)
    assert.NoError(t, err)
    assert.NotNil(t, validators.LoginValidator)
    assert.NotNil(t, validators.RegisterValidator)
}


func TestNewValidator_ErrorReadingSchema(t *testing.T) {
    mockReader := new(FileReaderMock)
    mockReader.On("ReadJsonSchema", "login_schema.json").Return([]byte{}, errors.New("error leyendo un esquema"))

    _, err := NewValidator(mockReader)
    assert.NotNil(t, err)

    mockReader.AssertExpectations(t)
}