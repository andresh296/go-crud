package user

import (
	"database/sql"
	"strings"

	domain "github.com/andresh296/go-crud/internal/domain/user"
)

const (
	queryGetAll = "SELECT id, name, age, email, password FROM users"
	querySave   = "INSERT INTO users (id, name, age, email, password) VALUES (?,?,?,?,?)"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.User, error) {
	stmt, err := r.db.Prepare(queryGetAll)
	if err != nil {
		return nil, domain.ErrGetUsers
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, domain.ErrUserCannotSave
	}
	defer rows.Close()

	usersDomain := []domain.User{}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password)
		if err != nil {
			return nil, domain.ErrUserCannotSave
		}

		usersDomain = append(usersDomain, user.ToDomain())
	}

	if err = rows.Err(); err != nil {
		return nil, domain.ErrUserCannotSave
	}

	return usersDomain, nil
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
			return domain.ErrNoUserSaved
		}
	}

	return nil
}
