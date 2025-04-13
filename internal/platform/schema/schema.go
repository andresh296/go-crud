package middleware

import (
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
		LoginValidator:   login,
		RegisterValidator: register,
	}, nil
}

func createSchema(resource string) (*jsonschema.Schema, error) {
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat = true
	schemaJSON, err := readJsonSchema(resource)
	if err != nil {
		return nil, err
	}
	schema, err := compiler.Compile(schemaJSON)
	if err != nil {
		return nil, err
	}  

	return schema, nil
}

func readJsonSchema(resource string)([]byte, error) {
	root, err := utils.FindModuleRoot()

	if err != nil {
		return nil, err
	}

	data, err := os.Open(filepath.Join(root, "internal/platform/schema/json_schemas",resource))

	if err != nil {
		return nil, err
	}

	return io.ReadAll(data)
}
	