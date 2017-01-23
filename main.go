package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"fmt"

	"./lib"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type Resource struct {
	Url []string
}

type Response struct {
	Users []lib.User
}

func main() {
	lib.SetPort("8080")
	lib.SwitchHandler(Actions())
}

func Actions() map[string]http.Handler {
	maps := make(map[string]http.Handler)

	maps["/hello"] = Hello{}
	maps["/insert"] = Insert{}
	maps["/select"] = Select{}

	return maps
}

type Hello struct {
}

func (a Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "select")
	lib.UserDao()
}

type Insert struct {
}

func (a Insert) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "insert")
	lib.InsertUserDao()
}

type Select struct {
}

func (a Select) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := Response{}
	res.Users = lib.SelectUserAllDao()

	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, json)
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
