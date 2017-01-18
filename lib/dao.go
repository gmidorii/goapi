package lib

import (
	"database/sql"
	"fmt"
	"log"
)

func UserDao() {
	db, err := sql.Open("mysql", "root@/hoge")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	table := "t_user"
	rows, err := db.Query(`
		SELECT * FROM ` + table)
	if err != nil {
		log.Fatal(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(columns)
}
