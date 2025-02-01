package user

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

// solo el molde
type MockDB struct {
	DB   *sql.DB
	Mock sqlmock.Sqlmock
}

