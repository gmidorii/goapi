package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	r, err := os.Open("./resource.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	//urls := readResource(r)

	lib.SetPort("8080")
	//lib.SetHandler(urls, ActionHandlers)
	lib.SwitchHandler(Actions())
}

//func ActionHandlers(w http.ResponseWriter, req *http.Request) {
//	u, err := url.Parse(req.RequestURI)
//	if err != nil {
//		log.Fatal(err)
//	}
//	param, err := url.ParseQuery(u.RawQuery)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// TODO: auto switch
//	switch u.Path {
//	case "/hello":
//		HelloAction(w, param)
//	case "/world":
//		fmt.Fprintln(w, "Hi")
//	case "/insert":
//		InsertAction(w, param)
//	case "/select":
//		SelectAction(w, param)
//	default:
//		fmt.Fprintln(w, "Default"+req.RequestURI)
//	}
//}

func Actions() map[string]interface {

	//actions := make([]lib.Handler, 0, 10)
	//
	//actions = append(actions, Hello{"/hello"})
	//actions = append(actions, Insert{"/insert"})
	//actions = append(actions, Select{"/select"})

	maps := map[string]interface{}

	maps["/hello"] = Hello{}
	maps["/insert"] = Insert{}
	maps["/select"] = Select{}

	return maps
}

type Hello struct {
}

func (a *Hello) HelloAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "select")
	lib.UserDao()
}

type Insert struct {
}

func (a *Insert) InsertAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "insert")
	lib.InsertUserDao()
}

type Select struct {
}

func (a *Select) SelectAction(w http.ResponseWriter, r *http.Request) {
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
