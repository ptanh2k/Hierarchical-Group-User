package group

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Group struct {
	Name         string `json:"name"`
	Name_in_path string
	Path         string `json:"path"`
}

func GetAllInfo(db *sql.DB, c *gin.Context) {
	rows, err := db.Query("SELECT * FROM group_;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("name\tname_in_path\tpath")

	for rows.Next() {
		var g Group
		if err := rows.Scan(&g.Name, &g.Name_in_path, &g.Path); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%s\t%s\n", g.Name, g.Name_in_path, g.Path)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}

func GetAllGroupPath(db *sql.DB) {
	rows, err := db.Query("SELECT path FROM group_;")

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
