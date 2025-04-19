package middleware

import (
	"encoding/json"
)

type MockGinContext struct {
	mockJSON   map[string]interface{}
	mockStatus int
	mockData   map[string]interface{}
}

func (m *MockGinContext) ToGinContext() any {
	panic("unimplemented")
}

type Context interface {
	ShouldBindJSON(obj interface{}) error
	JSON(code int, obj interface{})
	Set(key string, value interface{})
	Get(key string) (value interface{}, exists bool)
}

func (m *MockGinContext) ShouldBindJSON(obj interface{}) error {
	data, err := json.Marshal(m.mockJSON)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj)
}

func (m *MockGinContext) JSON(code int, obj interface{}) {
	m.mockStatus = code
	m.mockJSON = obj.(map[string]interface{})
}

func (m *MockGinContext) Set(key string, value interface{}) {
	m.mockData[key] = value
}

func (m *MockGinContext) Get(key string) (value interface{}, exists bool) {
	value, exists = m.mockData[key]
	return
}
