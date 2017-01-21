package lib

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Name  string `db:name`
	Color string `db:color`
}

var selectquery = `
SELECT name, color FROM t_user;
`

var insertquery = `
INSERT INTO t_user (name, color, update_time) VALUES ( ? , ? , now());
`

func UserDao() {
	db, err := sql.Open("mysql", "root@/hoge")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(selectquery)
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

func SelectUserAllDao() []User {
	db, err := sqlx.Connect("mysql", "root@/hoge")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	users := []User{}
	err = db.Select(&users, selectquery)
	if err != nil {
		log.Fatal(err)
	}

	for i, val := range users {
		fmt.Println(i)
		fmt.Println(val.Name)
		fmt.Println(val.Color)
	}

	return users
}

func InsertUserDao() {
	db, err := sql.Open("mysql", "root@/hoge")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	result, err := db.Exec(insertquery, "test", "red")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
