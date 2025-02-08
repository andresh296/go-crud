package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID string 
	Name string 
	Age int8 
	Email string 
	Password string 
}

func (u *User) setID() {
	u.ID = uuid.New().String()
}

func (u *User) hashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}
