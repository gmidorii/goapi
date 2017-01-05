package lib

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr *string

func SetPort(port string) {
	addr = flag.String("addr", ":"+port, "http service address")
}

func Handler(urls []string) {
	fmt.Println(urls)
	for _, url := range urls {
		http.Handle(url, http.HandlerFunc(allHandler))
	}
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func allHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "test")
}
