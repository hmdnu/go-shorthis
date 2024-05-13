package handlers

import (
	"net/http"

	"github.com/hmdnu/go-shorthis/common"
)

func HandleRedirect(w http.ResponseWriter, req *http.Request) {

	shortKey := req.PathValue("shortKey")

	if req.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if shortKey == "" {
		http.Error(w, "shortened key is missing", http.StatusBadRequest)
		return
	}

	originalUrl, found := common.Urls[shortKey]

	if !found {
		http.Error(w, "shortened key not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, req, originalUrl, http.StatusMovedPermanently)
}
