package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hmdnu/go-shorthis/config"
	_ "github.com/hmdnu/go-shorthis/database"
	"github.com/hmdnu/go-shorthis/handlers"
	"github.com/hmdnu/go-shorthis/middlewares"
)

func main() {

	http.HandleFunc("/{shortKey}", middlewares.Cors(handlers.HandleRedirect))
	http.HandleFunc("/", middlewares.Cors(handlers.HandleShorten))

	fmt.Println("server start in port " + config.PORT)
	if err := http.ListenAndServe(":"+config.PORT, nil); err != nil {
		log.Fatalln("cant listen to server", err.Error())
	}
}
