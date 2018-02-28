package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func connect() (*sqlx.DB, error) {
	fmt.Println("connect to DB")
	return sqlx.Connect("mysql", "root:password@(localhost:3306)/booking")
}
