package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/hmdnu/go-shorthis/model"
	"github.com/hmdnu/go-shorthis/utils"
)

func HandleRedirect(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		utils.ErrorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shortKey := req.PathValue("shortKey")

	if shortKey == "" {
		utils.ErrorResponse(w, "shortened key is missing", http.StatusBadRequest)
		return
	}

	// get from db
	var url model.URL

	if err := url.GetShortURL(shortKey); err != nil {

		if err == sql.ErrNoRows {
			utils.ErrorResponse(w, "cant find url shortcode", http.StatusNotFound)
			return
		}

		utils.ErrorResponse(w, "internal server error", http.StatusInternalServerError)
		log.Fatalln("error querying row", err.Error())
	}

	http.Redirect(w, req, url.OriginalUrl, http.StatusMovedPermanently)
}
