package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hmdnu/go-shorthis/config"
)

var DB *sql.DB

func init() {

	var err error

	DB, err = sql.Open("mysql", config.DATABASE)

	if err != nil {
		log.Fatalln("cant connect to database", err.Error())
	} else {
		fmt.Println("connected to database")
	}

}
