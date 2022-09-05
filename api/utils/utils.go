package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"cycir/group"
	"cycir/user"
)

func ConnectToDB(host string, port int, username string, password string, dbname string) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&group.Group{}, &user.User{})

	fmt.Println("Successfully connected")

	return db, nil
}
