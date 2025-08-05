package mysql

import (
	"database/sql"
	"fmt"
	"golang-codebase/src/configs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	configs.LoadEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		configs.GetEnv("DB_USER", "develop"),
		configs.GetEnv("DB_PASS", "develop123"),
		configs.GetEnv("DB_HOST", "localhost:3307"),
		configs.GetEnv("DB_NAME", "db_codebase"),
	)
	dba, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	dba.SetConnMaxLifetime(time.Minute * 3)
	dba.SetMaxOpenConns(10)
	dba.SetMaxIdleConns(10)
	db = dba
}

func GetDatabase() *sql.DB {
	return db
}
