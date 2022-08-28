package main

import (
	"cycir/group"
	"cycir/utils"
	"flag"
	"fmt"
)

type Config struct {
	Host	string	`mapstructure:"HOST"`
	Port	int32	`mapstructure:"PORT"`
	User	string	`mapstructure:"USER"`
	DB_name	string	`mapstructure:"DB_NAME"`
}


func main() {
	host := flag.String("host", "localhost", "Enter host (localhost): ")
	port := flag.Int("port", 5432, "Enter port (5432): ")
	user := flag.String("user", "postgres", "Enter user (postgres): ")
	password := flag.String("password", "password", "Enter password: ")
	dbname := flag.String("db", "database", "Enter database name: ")

	flag.Parse()

	db := utils.ConnectToDB(*host, *port, *user, *password, *dbname)

	defer db.Close()

	fmt.Println("Testing features: ")

	var choice int
	fmt.Print("Your choice? ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		group.GetAllGroup(db)
	default:
		break
	}
}