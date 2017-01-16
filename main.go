package main

import (
	"database/sql"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"fmt"

	"./lib"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type Resource struct {
	Url []string
}

func main() {
	r, err := os.Open("./resource.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	urls := readResource(r)

	lib.SetPort("8080")
	lib.SetHandler(urls, ActionHandler)
}

func ActionHandler(w http.ResponseWriter, req *http.Request) {
	u, err := url.Parse(req.RequestURI)
	if err != nil {
		log.Fatal(err)
	}
	param, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: auto switch
	switch u.Path {
	case "/hello":
		HelloAction(w, param)
	case "/world":
		fmt.Fprintln(w, "Hi")
	default:
		fmt.Fprintln(w, "Default"+req.RequestURI)
	}
}

func HelloAction(w http.ResponseWriter, param url.Values) {
	// sample
	json := `
	{
		"id": "test",
		"array": [
			"test1",
			"test2"
		]
	}
	`
	// response
	fmt.Fprintln(w, json)
	UserDao()
}

func UserDao() {
	db, err := sql.Open("mysql", "root:@/hoge")
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

	values := make([]sql.RawBytes, len(columns))

}

// readResource return resource file
func readResource(r io.Reader) []string {
	param, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	var resource Resource
	err = yaml.Unmarshal(param, &resource)
	if err != nil {
		log.Fatal(err)
	}
	return resource.Url
}
