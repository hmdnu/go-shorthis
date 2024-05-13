package main

import (
	"log"
	"net/http"

	"github.com/hmdnu/go-shorthis/config"
	"github.com/hmdnu/go-shorthis/handlers"
	"github.com/hmdnu/go-shorthis/middlewares"
)

func main() {

	http.HandleFunc("/", middlewares.Cors(handlers.HandleShorten))
	http.HandleFunc("/{shortKey}", middlewares.Cors(handlers.HandleRedirect))

	log.Println("server start in port " + config.PORT)
	if err := http.ListenAndServe(":"+config.PORT, nil); err != nil {
		log.Fatalln("cant listen to server", err.Error())
	}
}
