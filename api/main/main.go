package main

import (
	"cycir/group"
	"cycir/user"
	"cycir/utils"
	"flag"

	"github.com/gin-gonic/gin"
)

func main() {
	host := flag.String("host", "localhost", "Enter host (localhost): ")
	port := flag.Int("port", 5432, "Enter port (5432): ")
	username := flag.String("user", "postgres", "Enter user (postgres): ")
	password := flag.String("password", "password", "Enter password: ")
	dbname := flag.String("db", "database", "Enter database name: ")

	flag.Parse()

	db, err := utils.ConnectToDB(*host, *port, *username, *password, *dbname)

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	// Group
	router.GET("/groups", group.GetAllInfo(db))
	// router.GET("/groups/:cur_group_id/sub", group.GetSubGroup(db))
	router.POST("/groups", group.AddNewGroup(db))

	// User
	router.GET("/users", user.GetAllInfo(db))
	router.GET("/users/:id", user.GetUserById(db))

	router.Run("localhost:8080")
}
