package group

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Group struct {
	GID      int    `json:"gid"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id,omitempty"`
	Level    int    `json:"level,omitempty"`
	Path     string `json:"path,omitempty"`
}

// Get all info relate to group
func GetAllInfo(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		rows, err := db.Query("SELECT gid, name, parent_id level FROM group_;")

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close() // Should be called after handling error to avoid panic for trying to close a nil resultset

		groups := make([]Group, 0)

		for rows.Next() {
			var g Group
			if err := rows.Scan(&g.GID, &g.Name, &g.ParentID); err != nil {
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

		result, err := db.Exec("INSERT INTO group_ (gid, name, parent_id) VALUES ($1, $2, $3)",
			newGroup.GID, newGroup.Name, newGroup.ParentID)

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
		cur_group_id := c.Param("cur_group_id")

		rows, err := db.Query(`WITH RECURSIVE group_tree(gid, name, parent_id, lvl, gpath) AS (
									SELECT gid, name, parent_id, 1 AS level, name AS path
									FROM group_ 
									WHERE parent_id = $1
		  						UNION ALL
									SELECT bg.gid, bg.name, bg.parent_id, gt.lvl + 1, gt.gpath || ' - ' || bg.name
									FROM group_ bg
									JOIN group_tree gt ON bg.parent_id = gt.gid
								)
								SELECT * FROM group_tree;`, cur_group_id)

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		sub_groups := make([]Group, 0)

		for rows.Next() {
			var g Group
			if err := rows.Scan(&g.GID, &g.Name, &g.ParentID, &g.Level, &g.Path); err != nil {
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
