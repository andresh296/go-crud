package middleware

import (
	"github.com/stretchr/testify/mock"
)


type UtilsMock struct {
	mock.Mock
}

func (m *UtilsMock) FindModuleRoot() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}


type FileReaderMock struct {
	mock.Mock
}

func (m *FileReaderMock) ReadJsonSchema(resource string) ([]byte, error) {
    args := m.Called(resource)
    return args.Get(0).([]byte), args.Error(1)
}