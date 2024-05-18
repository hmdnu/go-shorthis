package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hmdnu/go-shorthis/config"
	"github.com/hmdnu/go-shorthis/model"
	"github.com/hmdnu/go-shorthis/utils"
)

type URL struct {
	Url      string `json:"url"`
	ShortUrl string `json:"shortUrl"`
}

func HandleShorten(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		utils.ErrorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var url URL

	err := json.NewDecoder(req.Body).Decode(&url)

	if err != nil {
		utils.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if url.Url == "" {
		utils.ErrorResponse(w, "url input not found", http.StatusNotFound)
		return
	}

	shortKey := utils.GenerateShortKey()
	uuid := uuid.NewString()

	// make struct to save to db
	shortUrl := model.URL{
		ID:          uuid,
		OriginalUrl: url.Url,
		ShortCode:   shortKey,
	}

	// save to db
	_, errDb := shortUrl.Insert()

	if errDb != nil {
		// if theres duplicate
		if errDb == model.ErrDuplicate {
			handleDuplicateURL(w, url.Url)
			return
		}

		utils.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
		log.Fatalln("cant insert into table url, error :", errDb.Error())
	}

	// assign short url to url struct
	url.ShortUrl = fmt.Sprintf("%s:%s/%s", config.SERVER_URL, config.PORT, shortKey)

	json, err := json.Marshal(&url)

	if err != nil {
		utils.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
		log.Fatalln("cant marshal json", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func handleDuplicateURL(w http.ResponseWriter, originalUrl string) {
	var url model.URL

	err := url.GetURL(originalUrl)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, "cant find url", http.StatusNotFound)
			return
		}

		utils.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
		log.Fatalln("error querying row", err.Error())
	}

	// write response
	var shortUrl URL

	shortUrl.Url = url.OriginalUrl
	shortUrl.ShortUrl = fmt.Sprintf("http://localhost:8080/%s", url.ShortCode)

	if b, err := json.Marshal(&shortUrl); err != nil {
		utils.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}

}
