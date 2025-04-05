package api

import (
	"fmt"

	domain "github.com/andresh296/go-crud/internal/domain/user"

	"github.com/go-playground/validator/v10"
)

type UserRequest struct {
	Name     string 
	Age      int8   
	Email    string 
	Password string 
}



type UserLogin struct {
	Email    string 
	Password string 
}

type LoginResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Email string `json:"email"`
	Token string `json:"token"`

}


type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Email string `json:"email"`
}

func (u UserRequest) ToDomain() domain.User {
	return domain.User{
		Name:     u.Name,
		Age:      u.Age,
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u UserLogin) ToDomain() domain.User {
	return domain.User{
		Email:    u.Email,
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

		return fmt.Errorf("%w: %s", ErrValidationUser, message)
	}
	return nil
}



func (u UserLogin) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		validateErrors := err.(validator.ValidationErrors)
		message := ""

		for _, validateErr := range validateErrors {
			message += fmt.Sprintf("%s: %s,", validateErr.Field(), validateErr.Error())
		}
		return fmt.Errorf(ErrValidationUser.Error(), message)
	}
	return nil
}
