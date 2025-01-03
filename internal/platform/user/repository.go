package user

import (
	"database/sql"
	"strings"

	domain "github.com/andresh296/go-crud/internal/domain/user"
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
		ID:       user.ID,
		Name:     user.Name,
		Age:      user.Age,
		Email:    user.Email,
		Password: user.Password,
	}

	stmt, err := r.db.Prepare(querySave)
	if err != nil {
		return domain.ErrUserCannotSave
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		userToSave.ID,
		userToSave.Name,
		userToSave.Age,
		userToSave.Email,
		userToSave.Password,
	)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Duplicate"):
			return domain.ErrDuplicateUser
		default:
			return domain.ErrUserCannotSave
		}
	}

	return nil
}
