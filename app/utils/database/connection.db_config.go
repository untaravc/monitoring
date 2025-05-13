package database

import (
	"database/sql"
	"fmt"
	"log"
	"monitoring/app/configs/db_config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMysqlDatabase() {
	var err error
	// Replace with your PostgreSQL connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME) // "%s:%s@tcp(%s:%s)

	// Open database connection
	fmt.Println(dsn)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}
	log.Println("✅ Connected to MySQL")
}

func InitPostgresDatabase() {
	var err error
	// Replace with your PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db_config.DB_HOST, db_config.DB_PORT, db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_NAME)

	// Open database connection
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection test failed:", err)
	}

	log.Println("Database connected successfully!")
}
