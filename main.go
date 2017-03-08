package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"fmt"

	"os"
	"strconv"
	"time"

	"./lib"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
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

	createCache()

	lib.SetPort("8080")
	lib.SwitchHandler(Actions())
}

func createCache() {
	for i := 0; i < 10000000; i++ {
		cache[strconv.Itoa(i)] = strconv.Itoa(i)
	}
}

func Actions() map[string]http.Handler {
	maps := make(map[string]http.Handler)

	maps["/hello"] = Hello{}
	maps["/insert"] = Insert{}
	maps["/select"] = Select{}
	maps["/load"] = Load{}

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

type Load struct {
}

func (a Load) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
