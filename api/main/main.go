package main

import (
	"cycir/group"
	"cycir/user"
	"cycir/utils"
	"flag"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Host    string `mapstructure:"HOST"`
	Port    int32  `mapstructure:"PORT"`
	User    string `mapstructure:"USER"`
	DB_name string `mapstructure:"DB_NAME"`
}

func main() {
	host := flag.String("host", "localhost", "Enter host (localhost): ")
	port := flag.Int("port", 5432, "Enter port (5432): ")
	username := flag.String("user", "postgres", "Enter user (postgres): ")
	password := flag.String("password", "password", "Enter password: ")
	dbname := flag.String("db", "database", "Enter database name: ")

	flag.Parse()

	db := utils.ConnectToDB(*host, *port, *username, *password, *dbname)

	defer db.Close()

	router := gin.Default()

	// Group
	router.GET("/groups", group.GetAllInfo(db))

	// User
	router.GET("/users", user.GetAllInfo(db))

	router.Run("localhost:8080")
}
