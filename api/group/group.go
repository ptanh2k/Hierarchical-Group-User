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
	Name_in_path string `json:"name_in_path,omitempty"`
	Path         string `json:"path,omitempty"`
	Parent       string `json:"parent,omitempty"`
	Level        int    `json:"level,omitempty"`
}

type Path struct {
	Path_name string `json:"path"`
	Parent    string `json:"parent"`
	Level     int    `json:"level"`
}

// Get all info relate to group
func GetAllInfo(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		rows, err := db.Query("SELECT name, path, subpath(path, 0, -1) parent, nlevel(path) level FROM group_;")

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close() // Should be called after handling error to avoid panic for trying to close a nil resultset

		groups := make([]Group, 0)

		for rows.Next() {
			var g Group
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

// Add new group to system
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

// Get all sub groups of a group and their level
func GetSubGroup(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		name := c.Param("name")

		rows, err := db.Query(`SELECT name, path, nlevel(path) level FROM group_ 
								WHERE path <@ (
									SELECT path FROM group_
									WHERE name = $1
								)
								AND path <> (
									SELECT path FROM group_
									WHERE name = $1
								)`, name)

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		sub_groups := make([]Group, 0)

		for rows.Next() {
			var g Group
			if err := rows.Scan(&g.Name, &g.Path, &g.Level); err != nil {
				log.Fatal(err)
			}
			sub_groups = append(sub_groups, g)
		}

		c.IndentedJSON(http.StatusOK, sub_groups)

		if err := rows.Err(); err != nil {
			panic(err)
		}
	}

	return gin.HandlerFunc(fn)
}
