package lib

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		// insert values via scanArgs pointer
		err := rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("----------------------------")
	}
}
