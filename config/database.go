package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found, using system environment variables.")
	}

	user := GetEnv("DB_USER", "root")
	password := GetEnv("DB_PASSWORD", "19072005")
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3306")
	dbname := GetEnv("DB_NAME", "backend_started")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		panic("Failed to connect to database!")
	}

	DB = database
	fmt.Println("Database connection successful!")
}
