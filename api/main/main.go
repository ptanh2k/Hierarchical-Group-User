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

	router.GET("/groups", group.GetAllInfo)
	router.GET("/users", user.GetAllInfo)

	// fmt.Println("Testing features: ")

	// var choice int
	// fmt.Print("Your choice? ")
	// fmt.Scanln(&choice)

	// switch choice {
	// case 1:
	// 	group.GetAllInfo(db)
	// case 2:
	// 	group.GetAllGroupPath(db)
	// case 3:
	// 	user.GetAllInfo(db)
	// default:
	// 	break
	// }
}
