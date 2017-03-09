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
	"gopkg.in/redis.v5"
	"gopkg.in/yaml.v2"
)

var cache = make(map[string]string)
var client *redis.Client

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
	insertCache()

	lib.SetPort("8080")
	lib.SwitchHandler(Actions())
}

func createCache() {
	for i := 0; i < 1000000; i++ {
		cache[strconv.Itoa(i)] = strconv.Itoa(i)
	}
}

func insertCache() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	for k, v := range cache {
		err := client.Set(k, v, 0).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Actions() map[string]http.Handler {
	maps := make(map[string]http.Handler)

	maps["/hello"] = Hello{}
	maps["/insert"] = Insert{}
	maps["/select"] = Select{}
	maps["/load"] = Load{}
	maps["/load-redis"] = LoadRedis{}

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

type LoadRedis struct{}

func (a LoadRedis) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	v1, _ := client.Get("318419").Result()
	fmt.Println(v1)
	v2, _ := client.Get("71897").Result()
	fmt.Println(v2)
	v3, _ := client.Get("52987").Result()
	fmt.Println(v3)
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
