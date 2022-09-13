package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	host     = "localhost"
	port     = 3306
	username = "root"
	password = "Erickype"
	dbname   = "test"
)

func GetMysqlConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, err
}
