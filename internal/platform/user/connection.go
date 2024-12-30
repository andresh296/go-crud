package user

import (
	"github.com/andresh296/go-crud/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB(config config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Schema)
	db, err := sql.Open(config.Driver, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}