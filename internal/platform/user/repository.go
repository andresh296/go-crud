package user

import (
	domain "github.com/andresh296/go-crud/internal/domain/user"
	"database/sql"
)

const (
	queryGetAll = "SELECT id, name, age, email, password FROM users"
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

func (r *repository) GetAll() ([]domain.User, error) {
	stmt, err := r.db.Prepare(queryGetAll)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	usersDomain := []domain.User{}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}

		usersDomain = append(usersDomain, user.ToDomain())
	}

	return usersDomain, nil
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