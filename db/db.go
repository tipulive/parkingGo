package db

import (
	"database/sql"
	"fmt"
	"log"

	"parking/config" // replace with your actual module path

	_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
)

var DB *sql.DB

//var DB *sqlx.DB

func InitDB(cfg config.MySQLConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	//DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	fmt.Println("Database connected!")
}
