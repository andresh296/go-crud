package middleware

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/andresh296/go-crud/tools/utils"
	"github.com/kaptinlin/jsonschema"
)

type Validators struct {
	LoginValidator *jsonschema.Schema
	RegisterValidator *jsonschema.Schema
}

type FileReaderInterface interface {
	ReadJsonSchema(resource string) ([]byte, error)
}

type DefaultFileReader struct{}

var fileReader FileReaderInterface = &DefaultFileReader{}

func SetFileReader(reader FileReaderInterface) {
	fileReader = reader
}



func NewValidator() (*Validators, error) {
	login, err := createSchema("login_schema.json")
	if err != nil {
		return nil, err
	}

	register, err := createSchema("register_schema.json")
	if err != nil {
		return nil, err
	}

	return &Validators{
		LoginValidator:    login,
		RegisterValidator: register,
	}, nil
}

func createSchema(resource string) (*jsonschema.Schema, error) {
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat = true
	schemaJSON, err := fileReader.ReadJsonSchema(resource)
	if err != nil {
		return nil, errors.New("failed to read JSON schema: " + err.Error())
	}
	if schemaJSON == nil {
		return nil, errors.New("schemaJSON is nil or empty")
	}
    schema, err := compiler.Compile(schemaJSON)
    if err != nil {
        return nil, err
    }

    return schema, nil
}


func (r *DefaultFileReader) ReadJsonSchema(resource string) ([]byte, error) {
	root, err := utils.FindModuleRoot()
	if err != nil {
		return nil, err
	}

	data, err := os.Open(filepath.Join(root, "internal/platform/schema/json_schemas", resource))
	if err != nil {
		return nil, err
	}
	defer data.Close()

	return io.ReadAll(data)
}