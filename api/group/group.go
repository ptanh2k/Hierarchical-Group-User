package group

import (
	"database/sql"
	"fmt"
	"log"
)

type Group struct {
	Name string
	Name_in_path string
	Path string
}

func GetAllGroup(db *sql.DB) {
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
