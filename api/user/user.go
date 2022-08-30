package user

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Uid       int    `json:"uid,omitempty"`
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
			panic(err)
		}

		defer rows.Close()

		users := make([]User, 0)

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

// Get user by user ID
func GetUserById(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")

		stmt, err := db.Prepare(`SELECT u.username, u.firstname, u.lastname, u.email, g.path 
								FROM user_ u INNER JOIN group_ g 
								ON u.gid = g.gid WHERE u.uid = $1;`)

		if err != nil {
			panic(err)
		}

		var u User

		err = stmt.QueryRow(id).Scan(&u.Username, &u.Firstname, &u.Lastname, &u.Email, &u.Path)

		if err != nil {
			if err == sql.ErrNoRows {
				c.IndentedJSON(http.StatusNotFound, u)
			}
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, u)
	}

	return gin.HandlerFunc(fn)
}
