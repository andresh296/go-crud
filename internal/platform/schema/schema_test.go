package middleware

import (
	//"errors"
	"errors"
	"testing"

	"github.com/test-go/testify/assert"
	"github.com/test-go/testify/mock"
)

func TestNewValidator_MockConfiguration(t *testing.T) {
 
    mockReader := new(FileReaderMock)
    mockReader.On("ReadJsonSchema", "login_schema.json").Return([]byte(`{}`), nil)
    mockReader.On("ReadJsonSchema", "register_schema.json").Return([]byte(`{}`), nil)

    originalReader := fileReader
    SetFileReader(mockReader)
    defer SetFileReader(originalReader)
    _, _ = NewValidator()
   
    mockReader.AssertExpectations(t)
}

func TestNewValidator(t *testing.T) {
    
    originalReader := fileReader

    
	mockReader := new(FileReaderMock)
	mockReader.On("ReadJsonSchema", "login_schema.json").Return([]byte(`{}`), nil)
	mockReader.On("ReadJsonSchema", "register_schema.json").Return([]byte(`{}`), nil)

	
	SetFileReader(mockReader)
	defer SetFileReader(originalReader)

    validators, err := NewValidator()
    assert.NoError(t, err)
    assert.NotNil(t, validators.LoginValidator)
    assert.NotNil(t, validators.RegisterValidator)

}



func TestNewValidator_ErrorReadingSchema(t *testing.T) {
    // Configurar el mock para devolver un error
    mockReader := new(FileReaderMock)
    mockReader.On("ReadJsonSchema", mock.Anything).Return([]byte{}, errors.New("error leyendo un esquema"))


    // Reemplazar la variable fileReader con el mock
    originalReader := fileReader
    fileReader = mockReader
    defer func() {
        fileReader = originalReader
    }()

    // Testear que se devuelva un error si no se puede leer el esquema de login
    _, err := NewValidator()
    assert.NotNil(t, err)

    // Verificar que el mock haya sido llamado correctamente
    mockReader.AssertExpectations(t)
}