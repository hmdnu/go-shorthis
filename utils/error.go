package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hmdnu/go-shorthis/common"
)

func ErrorResponse(w http.ResponseWriter, message string, status int) {

	w.WriteHeader(status)

	res := common.Error{
		Message: message,
		Status:  status,
	}

	b, err := json.Marshal(&res)

	if err != nil {
		log.Fatalln("cant marshall json", err.Error())
	}

	w.Write(b)
}
