package model

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/hmdnu/go-shorthis/database"
)

var (
	ErrDuplicate = errors.New("duplicate")
)

type URL struct {
	ID          string `json:"id"`
	OriginalUrl string `json:"url"`
	ShortCode   string `json:"shortUrl"`
}

func (url *URL) Insert() (sql.Result, error) {
	res, err := database.DB.Exec("INSERT INTO url VALUES(?, ?, ?);", &url.ID, &url.OriginalUrl, &url.ShortCode)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return nil, ErrDuplicate
		}
		return nil, err
	}

	return res, nil

}

func (url *URL) GetShortURL(shortKey string) error {

	err := database.DB.QueryRow("SELECT original_url, short_code FROM url WHERE short_code = ?", shortKey).
		Scan(&url.OriginalUrl, &url.ShortCode)

	return err
}

func (url *URL) GetURL(originalUrl string) error {
	err := database.DB.QueryRow("SELECT original_url, short_code FROM url WHERE original_url = ?", originalUrl).Scan(&url.OriginalUrl, &url.ShortCode)

	return err
}
