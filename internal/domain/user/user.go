package user

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/google/uuid"
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

func (u *User) hashPassword() {
	hasher := sha256.New()
	hasher.Write([]byte(u.Password))
	u.Password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
