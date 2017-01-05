package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

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
	lib.Handler(urls)
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
