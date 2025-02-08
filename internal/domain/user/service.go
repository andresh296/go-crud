package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	GetByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	Save(user User) error
}

type Service interface {
	GetByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	Save(user User) (User, error)
	Login(user User) (*User, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s service) GetByID(id string) (*User, error) {
	return s.repository.GetByID(id)
}

func (s service) GetUserByEmail(email string) (*User, error) {
	return s.repository.GetUserByEmail(email)
}

func (s service) Save(user User) (User, error) {
	user.setID()
	user.hashPassword()
	err := s.repository.Save(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u User) comparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

func (s service) Login(user User) (*User, error) {
	userFound, err := s.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if err := userFound.comparePassword(user.Password); err != nil {
		return nil, err
	}

	return userFound, nil
}
