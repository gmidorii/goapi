package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"net/http"

	"fmt"

	"./lib"
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
	switch req.RequestURI {
	case "/hello":
		fmt.Fprintln(w, req.RequestURI)
	case "/world":
		fmt.Fprintln(w, "Hi")
	}
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
