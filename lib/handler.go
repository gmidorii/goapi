package lib

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr *string

type handlerFunc func(w http.ResponseWriter, req *http.Request)

func SetPort(port string) {
	addr = flag.String("addr", ":"+port, "http service address")
}

func SetHandler(urls []string, fn handlerFunc) {
	fmt.Println(urls)
	for _, url := range urls {
		http.Handle(url, http.HandlerFunc(fn))
	}
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
