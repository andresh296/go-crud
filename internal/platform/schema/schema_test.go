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
    // Configurar el mock para devolver un error
    mockReader := new(FileReaderMock)
    mockReader.On("ReadJsonSchema", "login_schema.json").Return([]byte{}, errors.New("error leyendo un esquema"))

    // Testear que se devuelva un error si no se puede leer el esquema de login
    _, err := NewValidator(mockReader)
    assert.NotNil(t, err)

    // Verificar que el mock haya sido llamado correctamente
    mockReader.AssertExpectations(t)
}