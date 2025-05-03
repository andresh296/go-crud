package middleware

import (
	"github.com/stretchr/testify/mock"
)


type FileReaderMock struct {
	mock.Mock
}

func (m *FileReaderMock) ReadJsonSchema(resource string) ([]byte, error) {
    args := m.Called(resource)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
    return args.Get(0).([]byte), args.Error(1)
}