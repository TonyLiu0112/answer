package web

import (
	"com.tony666.answer/web/handlers"
	"log"
	"net/http"
)

func Init() {
	log.Println("do init routes.")
	http.HandleFunc("/answer", handlers.DoAnswer)
	http.HandleFunc("/begin", handlers.BeginActivity)
}
