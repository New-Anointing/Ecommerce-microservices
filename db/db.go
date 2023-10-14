package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	var (
		dbUsername = os.Getenv("DB_USERNAME")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbHost     = os.Getenv("DB_HOST")
		dbName     = os.Getenv("DB_DBNAME")
		dbPortStr  = os.Getenv("DB_PORT")
	)
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to port: %v", err))
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Could not establish db connection: %v", err))
	}
	if gin.Mode() == gin.ReleaseMode {
		db.Logger.LogMode(0)
	}
	DB = db
	rawDB := RawDB()
	rawDB.SetMaxIdleConns(20)
	rawDB.SetMaxOpenConns(100)

	err = migrate()
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate DB: %v", err))

	}

}

func RawDB() *sql.DB {
	db, err := DB.DB()
	if err != nil {
		panic(err)
	}

	return db
}
