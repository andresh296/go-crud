package user

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockDB struct {
	DB   *sql.DB
	Mock sqlmock.Sqlmock
}

