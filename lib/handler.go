package lib

import (
	"flag"
	"log"
	"net/http"
)

var addr *string

func SetPort(port string) {
	addr = flag.String("addr", ":"+port, "http service address")
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
