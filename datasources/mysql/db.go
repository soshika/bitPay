package mysql

import (
	"bitPay/logger"
	"database/sql"
	"fmt"
	"os"
)

var (
	Client *sql.DB
)


func init() {
	args 		:= os.Args
	jsonPath 	:= ""

	if len(args) < 2 {
		jsonPath = "assets/json/user_db.json"
	}else {
		jsonPath = args[1]
	}

	infoDB, parseErr := ParseJson(jsonPath)
	if parseErr != nil {
		logger.Error("error when tyring to parse json", parseErr)
		panic(parseErr)
	}

	username := infoDB["username"]
	password := infoDB["password"]
	host 	 := infoDB["host"]
	schema   := infoDB["schema"]

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

