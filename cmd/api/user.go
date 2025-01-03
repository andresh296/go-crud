package api

import (
	domain "github.com/andresh296/go-crud/internal/domain/user"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UserRequest struct {
	Name string `json:"name" validate:"required,max=100"`
	Age int8 `json:"age" validate:"required,gte=18"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (u UserRequest) ToDomain() domain.User {
	return domain.User{
		Name: u.Name,
		Age: u.Age,
		Email: u.Email,
		Password: u.Password,
	}
}

func (u UserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		validateErrors := err.(validator.ValidationErrors)
		message := ""
		
		for _, validateErr := range validateErrors {
			message += fmt.Sprintf("%s: %s,", validateErr.Field(), validateErr.Error())
		}

		return fmt.Errorf(ErrValidationUser.Error(),message)

		}
		return nil
	}


type UserResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Age int8 `json:"age"`
	Email string `json:"email"`
}
