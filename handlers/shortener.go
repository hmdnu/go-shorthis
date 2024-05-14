package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hmdnu/go-shorthis/config"
	"github.com/hmdnu/go-shorthis/model"
	"github.com/hmdnu/go-shorthis/utils"
)

func HandleShorten(w http.ResponseWriter, req *http.Request) {
	var URL model.URL

	if req.Method != http.MethodPost {
		utils.ErrorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(req.Body).Decode(&URL)

	if err != nil {
		utils.ErrorResponse(w, "internal server error", http.StatusMethodNotAllowed)
		return
	}

	if URL.OriginalUrl == "" {
		utils.ErrorResponse(w, "url input not found", http.StatusNotFound)
		return
	}

	shortKey := utils.GenerateShortKey()
	uuid := utils.GenerateUUID()

	// make struct to save to db
	url := model.URL{
		ID:          uuid,
		OriginalUrl: URL.OriginalUrl,
		ShortCode:   shortKey,
	}

	// save to db
	if _, err := url.Insert(); err != nil {
		utils.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
		log.Fatalln("cant insert into table url, error :", err.Error())
	}

	// assign URL struct
	URL.ID = uuid
	URL.ShortCode = fmt.Sprintf("%s:%s/%s", config.SERVER_URL, config.PORT, shortKey)

	json, err := json.Marshal(&URL)

	if err != nil {
		utils.ErrorResponse(w, "cant marshal json", http.StatusInternalServerError)
		log.Fatalln("cant marshal json", err.Error())
	}

	w.Write(json)
}
