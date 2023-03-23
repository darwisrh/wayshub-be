package mysql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection Database
func DatabaseInit() {
	var err error
	var DB_HOST = "containers-us-west-132.railway.app"
	var DB_USER = "postgres"
	var DB_PASSWORD = "17SwerQUGezUe1JymOAw"
	var DB_NAME = "railway"
	var DB_PORT = "5582"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
