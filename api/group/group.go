package group

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Group struct {
	Name         string `json:"name"`
	Name_in_path string `json:"name_in_path"`
	Path         string `json:"path"`
}

type Path struct {
	Path_name string `json:"path"`
	Parent    string `json:"parent"`
	Level     int    `json:"level"`
}

func GetAllInfo(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		rows, err := db.Query("SELECT name, path, subpath(path, 0, -1) parent, nlevel(path) level FROM group_;")

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close() // Should be called after handling error to avoid panic for trying to close a nil resultset

		type GroupWithLevel struct {
			Name   string `json:"name"`
			Path   string `json:"path"`
			Parent string `json:"parent"`
			Level  int    `json:"level"`
		}

		groups := make([]GroupWithLevel, 0)

		for rows.Next() {
			var g GroupWithLevel
			if err := rows.Scan(&g.Name, &g.Path, &g.Parent, &g.Level); err != nil {
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

func AddNewGroup(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var newGroup Group

		// Bind the received JSON to newGroup
		if err := c.BindJSON(&newGroup); err != nil {
			return
		}

		result, err := db.Exec("INSERT INTO group_ (name, name_in_path, path) VALUES ($1, $2, $3)",
			newGroup.Name, newGroup.Name_in_path, newGroup.Path)

		if err != nil {
			panic(err)
		}

		rowsAffected, err := result.RowsAffected()

		if err != nil {
			panic(err)
		}

		fmt.Printf("Created successfully (%d row(s) affected)", rowsAffected)
	}

	return gin.HandlerFunc(fn)
}
