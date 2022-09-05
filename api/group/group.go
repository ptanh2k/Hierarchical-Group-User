package group

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Group struct {
	GID      int    `json:"gid" gorm:"primaryKey"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
}

type Tabler interface {
	TableName() string
}

// Override default table name
func (Group) TableName() string {
	return "group_"
}

// Get all info relate to group
func GetAllInfo(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		groups := make([]Group, 0)

		db.Table("group_").Find(&groups)

		c.IndentedJSON(http.StatusOK, gin.H{"groups": groups})

	}
	return gin.HandlerFunc(fn)
}

// Add new group to system
func AddNewGroup(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		type CreateGroupInput struct {
			GID      int    `json:"gid" binding:"required"`
			Name     string `json:"name" binding:"required"`
			ParentID int    `json:"parent_id" binding:"required"`
		}

		var input CreateGroupInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newGroup := Group{GID: input.GID, Name: input.Name, ParentID: input.ParentID}

		result := db.Create(&newGroup)

		if err := result.Error; err != nil {
			panic(err)
		}

		fmt.Printf("%d row(s) affected", result.RowsAffected)

		c.JSON(http.StatusCreated, gin.H{"group": newGroup})
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
			if err := rows.Scan(&g.GID, &g.Name, &g.ParentID); err != nil {
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
