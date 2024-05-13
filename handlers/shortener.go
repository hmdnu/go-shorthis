package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hmdnu/go-shorthis/common"
	"github.com/hmdnu/go-shorthis/config"
	"github.com/hmdnu/go-shorthis/utils"
)

func HandleShorten(w http.ResponseWriter, req *http.Request) {
	var URL common.URL

	if req.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(req.Body).Decode(&URL)

	if err != nil {
		http.Error(w, "cant decode json", http.StatusInternalServerError)
		return
	}

	if URL.OriginalUrl == "" {
		http.Error(w, "url input not found", http.StatusNotFound)
		return
	}

	shortKey := utils.GenerateShortKey()
	common.Urls[shortKey] = URL.OriginalUrl
	URL.ShortenUrl = fmt.Sprintf("%s:%s/%s", config.SERVER_URL, config.PORT, shortKey)

	json, err := json.Marshal(URL)

	if err != nil {
		http.Error(w, "cant marshal json", http.StatusInternalServerError)
		log.Fatalln("cant marshal json", err.Error())
	}

	w.Write(json)
}
