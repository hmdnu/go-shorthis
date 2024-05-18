package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func ErrorResponse(w http.ResponseWriter, message string, status int) {

	w.WriteHeader(status)

	res := Error{
		Message: message,
		Status:  status,
	}

	b, err := json.Marshal(&res)

	if err != nil {
		log.Fatalln("cant marshall json", err.Error())
	}

	w.Write(b)
}
