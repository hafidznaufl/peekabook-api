package config

import (
	"fmt"
	"os"
	"peekabook/model/schema"

	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		panic("Failed to Connect Database")
	}

	Migrate()

	fmt.Println("Connected to Database")

	return DB
}

func Migrate() {
	DB.AutoMigrate(schema.User{}, schema.Admin{}, schema.Author{}, schema.Book{}, schema.Chat{}, schema.Store{}, schema.Borrow{})
}
