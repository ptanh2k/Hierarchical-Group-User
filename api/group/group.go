package group

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Group struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Level int    `json:"level"`
}

func GetAllInfo(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		rows, err := db.Query("SELECT name, path, nlevel(path) level FROM group_;")

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()		// Should be called after handling error to avoid panic for trying to close a nil resultset

		var groups []Group

		for rows.Next() {
			var g Group
			if err := rows.Scan(&g.Name, &g.Path, &g.Level); err != nil {
				log.Fatal(err)
			}
			groups = append(groups, g)
		}
		c.IndentedJSON(http.StatusOK, groups)

		if err := rows.Err(); err != nil {
			panic(err)
		}
	}
	return gin.HandlerFunc(fn)
}

func GetAllGroupPath(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		rows, err := db.Query("SELECT path, subpath(path, 0, -1), nlevel(path) as parent FROM group_;")

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			var path string
			if err := rows.Scan(&path); err != nil {
				log.Fatal(err)
			}
			fmt.Println(path)
		}

		if err := rows.Err(); err != nil {
			panic(err)
		}
	}
	return gin.HandlerFunc(fn)
}
