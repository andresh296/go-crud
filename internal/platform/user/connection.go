package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andresh296/go-crud/config"

	_ "github.com/go-sql-driver/mysql"

)

func GetDB(config config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Schema)
	db, err := sql.Open(config.Driver, dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}