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

func SetHandler(urls []string, fn http.HandlerFunc) {
	fmt.Println(urls)
	for _, url := range urls {
		http.Handle(url, fn)
	}
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func SwitchHandler(hadlerMap map[string]http.Handler) {
	for key, val := range hadlerMap {
		http.Handle(key, val)
	}
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
