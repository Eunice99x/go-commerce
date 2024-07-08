package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB

func Connect() error {
	dataSourceName := "root:root@tcp(mysql:3306)/mydatabase"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	boil.SetDB(db)
	DB = db

	return nil
}
