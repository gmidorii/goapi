package main

import (
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"

	"fmt"

	"./lib"
	"gopkg.in/yaml.v2"
)

var tmpl = template.Must(template.New("").Parse(`package {{.Package}}

func {{.Path}}Handler(w http.ResponseWriter, req *http.Request) {
	{{.Path}}Action(w, req)
}

`))

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

	if err := Generate("hello"); err != nil {
		log.Fatal(err)
	}
	lib.SetPort("8080")
	lib.SetHandler(urls, helloHandler)
}

func Generate(path string) error {
	pkg, err := build.Default.ImportDir(".", 0)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(fmt.Sprintf("%s_gen.go", strings.ToLower(path)))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	return tmpl.Execute(f, map[string]interface{}{
		"Package": pkg.Name,
		"Path":    path,
	})
}

func HelloAction(w http.ResponseWriter, req *http.Request) {
	json := `
	{
		"id": "test",
		"array": [
			"test10",
			"test20"
		]
	}
	`
	// response
	fmt.Fprintln(w, json)
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
		HellosAction(w, param)
	case "/world":
		fmt.Fprintln(w, "Hi")
	default:
		fmt.Fprintln(w, "Default"+req.RequestURI)
	}
}

func HellosAction(w http.ResponseWriter, param url.Values) {
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
