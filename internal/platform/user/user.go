package user

import domain "github.com/andresh296/go-crud/internal/domain/user"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) ToDomain() domain.User {
	return domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Age:      u.Age,
		Email:    u.Email,
		Password: u.Password,
	}
}
