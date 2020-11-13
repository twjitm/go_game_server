package database

import (
	"database/sql"
)

func GetDriver() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/caiop?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}

