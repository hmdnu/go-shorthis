package model

import (
	"database/sql"

	"github.com/hmdnu/go-shorthis/database"
)

type URL struct {
	ID          string `json:"id"`
	OriginalUrl string `json:"originalUrl"`
	ShortCode   string `json:"shortCode"`
}

func (url *URL) Insert() (sql.Result, error) {
	res, err := database.DB.Exec("INSERT INTO url VALUES(?, ?, ?);", &url.ID, &url.OriginalUrl, &url.ShortCode)

	return res, err

}

func (url *URL) GetShortCode(shortKey string) error {

	err := database.DB.QueryRow("SELECT * FROM url WHERE short_code = ?", shortKey).
		Scan(&url.ID, &url.OriginalUrl, &url.ShortCode)

	return err
}
