package model

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)

func connect() (*sqlx.DB, error) {
    return sqlx.Connect("mysql", "root:password@(localhost:3306)/booking")
}
