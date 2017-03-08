package main

import (
	"encoding/json"
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
	"strconv"
	"time"
)

var cache = make(map[string]string)

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
	urls := readResource(r)

	createCache()

	lib.SetPort("8080")
	lib.SetHandler(urls, ActionHandler)
}

func createCache() {
	for i := 0; i < 10000000; i++ {
		cache[strconv.Itoa(i)] = strconv.Itoa(i)
	}
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
	case "/insert":
		InsertAction(w, param)
	case "/select":
		SelectAction(w, param)
	case "/load":
		load(w)
	default:
		fmt.Fprintln(w, "Default"+req.RequestURI)
	}
}

func HelloAction(w http.ResponseWriter, param url.Values) {
	fmt.Fprintln(w, "select")
	lib.UserDao()
}

func InsertAction(w http.ResponseWriter, param url.Values) {
	fmt.Fprintln(w, "insert")
	lib.InsertUserDao()
}

func SelectAction(w http.ResponseWriter, param url.Values) {
	res := Response{}
	res.Users = lib.SelectUserAllDao()

	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, json)
}

func load(w http.ResponseWriter) {
	start := time.Now()
	fmt.Println(cache["318419"])
	fmt.Println(cache["71897"])
	fmt.Println(cache["52987"])
	end := time.Now()
	fmt.Fprintf(w, "%f\n", end.Sub(start).Seconds())
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
