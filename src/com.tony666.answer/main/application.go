package main

import (
	"com.tony666.answer/web"
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	web.Init()
	log.Fatal(http.ListenAndServe(*addr, nil))
}
