package mysql

import (
	"bitPay/logger"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	Client *sql.DB
)


func init() {
	args 		:= os.Args
	jsonPath 	:= ""

	if len(args) < 2 {
		jsonPath = "assets/json/db.json"
	}else {
		jsonPath = args[1]
	}

	_, parseErr := ParseJson(jsonPath)
	if parseErr != nil {
		logger.Error("error when tyring to parse json", parseErr)
		panic(parseErr)
	}

	// TODO: should change for product or add environment variable
	username := os.Getenv("username")
	password := os.Getenv("password")
	host 	 := os.Getenv("host")
	schema   := os.Getenv("schema")

	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	logger.Info("database successfully configured!")
}

