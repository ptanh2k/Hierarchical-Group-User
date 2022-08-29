package user

import (
	"database/sql"
	"log"
	"net/http"

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

func GetAllInfo(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM user_;")

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		var users []User

		for rows.Next() {
			var u User
			if err := rows.Scan(&u.Uid, &u.Username, &u.Firstname, &u.Lastname, &u.Email, &u.Path); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}

		c.IndentedJSON(http.StatusOK, users)

		if err := rows.Err(); err != nil {
			panic(err)
		}
	}
	return gin.HandlerFunc(fn)
}
