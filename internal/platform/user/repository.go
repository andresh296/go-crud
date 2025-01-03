package user

import (
	domain "github.com/andresh296/go-crud/internal/domain/user"
	"database/sql"
)

const (
	querySave = "INSERT INTO users (id, name, age, email, password) VALUES (?,?,?,?,?)"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}
}


func (r repository) Save(user domain.User) error {
	userToSave := User{
		ID: user.ID,
		Name: user.Name,
		Age: user.Age,
		Email: user.Email,
		Password: user.Password,
	}

	stmt, err := r.db.Prepare(querySave)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		userToSave.ID,
		userToSave.Name,
		userToSave.Age,
		userToSave.Email,
		userToSave.Password,
	)

	return err
}