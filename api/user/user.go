package user

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	uid       int
	username  string
	firstname string
	lastname  string
	email     string
	path      string
}

func GetAllInfo(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM user_;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("username\tfirstname\tlastname\tpath")

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.uid, &u.username, &u.firstname, &u.lastname, &u.email, &u.path); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%s\t%s\t%s\n", u.username, u.firstname, u.lastname, u.path)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
