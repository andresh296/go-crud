package user

import "errors"

var (
	ErrUserCannotSave = errors.New("error user can not save")
	ErrGetUsers       = errors.New("error get users")
	ErrDuplicateUser  = errors.New("user already exists")
	ErrSavingUser     = errors.New("error saving user")
	ErrGettingEmail    =errors.New("error getting the email")
)
