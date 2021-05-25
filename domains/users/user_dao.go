package users

import (
	"bitPay/datasources/mysql"
	"bitPay/logger"
	"bitPay/utills/responses"
	"net/http"
)

const (
	queryInsertUserAuth		= "INSERT INTO user (id, api_key, secret_key, is_test) VALUES(?, ?, ?, ?);"
	querySelectUser			= "SELECT id, secret_key FROM user WHERE api_key = ?;"
)

func (user *User) Save() *responses.Response {
	stmt, err := mysql.Client.Prepare(queryInsertUserAuth)
	if err != nil {
		logger.Error("error when trying to prepare save User statement.", err)
		return responses.NewInternalServerError("Internal Server Error", "Please Try Again Later...")
	}
	defer stmt.Close()

	result, saveErr := stmt.Exec(user.APIKey, user.SecretKey, user.IsTest)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return responses.NewBadRequestError("User Registered before!", "Please make sure to enter new api_key", http.StatusBadRequest)
	}

	user.Id, _ = result.LastInsertId()

	return nil
}

func (user *User) Get() *responses.Response {
	stmt, err := mysql.Client.Prepare(querySelectUser)
	if err != nil {
		logger.Error("error when trying to prepare get user by api_key statement", err)
		return responses.NewInternalServerError("Internal Server Error", "Please Try Again later...")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.APIKey)
	if getErr := result.Scan(&user.Id, &user.SecretKey); getErr != nil {
		logger.Error("error when trying to get user by api-key", getErr)
		return responses.NewInternalServerError("Internal Server Error", "Please Try Again later...")
	}

	return nil
}