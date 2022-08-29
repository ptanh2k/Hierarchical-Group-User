package user

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	Uid       int    `json:"uid"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Path      string `json:"path"`
}

func GetAllInfo(db *sql.DB, c *gin.Context) {
	rows, err := db.Query("SELECT * FROM user_;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("username\tfirstname\tlastname\tpath")

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Uid, &u.Username, &u.Firstname, &u.Lastname, &u.Email, &u.Path); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%s\t%s\t%s\n", u.Username, u.Firstname, u.Lastname, u.Path)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
