package lib

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr *string

//type handlerFunc func(w http.ResponseWriter, req *http.Request)

//type Handler interface {
//	Handler(w http.ResponseWriter, req *http.Request)
//}

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

func SwitchHandler(hmap map[string]Handler) {
	for key, val := range hmap {
		http.Handle(key, val.Handler)
	}
}
