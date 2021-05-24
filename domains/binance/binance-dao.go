package binance

import (
	"bitPay/datasources/mysql"
	"bitPay/logger"
	"bitPay/utills/responses"
)

const (
	queryInsertOrder = "INSERT INTO order (user_id, symbol, order_id, client_order_id, transact_time, price) VALUES(?, ?, ?, ?, ?, ?);"
	querySelectOrder = "SELECT id, symbol, client_order_id, transact_time, price FROM order WHERE user_id=? AND order_id=?;"
	queryDeleteOrder = "DELETE FROM order WHERE user_id=? AND order_id=?;"
)

func (order *Order) Save() *responses.Response {
	stmt, err := mysql.Client.Prepare(queryInsertOrder)
	if err != nil {
		logger.Error("error when trying to prepare Save order statement!", err)
		return responses.NewInternalServerError("Internal Server Error", "Please try again later...")
	}
	defer stmt.Close()

	result, saveErr := stmt.Exec(order.UserId, order.Symbol, order.OrderID, order.ClientOrderID, order.TransactTime, order.Price)
	if saveErr != nil {
		logger.Error("error when trying to save order statement", saveErr)
		return responses.NewInternalServerError("Internal Server Error", "Please try again later...")

	}

	order.Id, _ = result.LastInsertId()

	return nil
}

func (order *Order) Get() *responses.Response {
	stmt, err := mysql.Client.Prepare(querySelectOrder)
	if err != nil {
		logger.Error("error when trying to prepare get order statement", err)
		return responses.NewInternalServerError("Internal Server Error", "Please try again later...")
	}
	defer stmt.Close()

	result := stmt.QueryRow(order.UserId, order.OrderID)
	if getErr := result.Scan(&order.Id, &order.Symbol, &order.ClientOrderID, &order.TransactTime, &order.Price) ; getErr != nil {
		logger.Error("error when trying to get hotel-detail statement", getErr)
		return responses.NewInternalServerError("Internal Server Error", "")

	}

	return nil
}

func (order *Order) Delete() *responses.Response {
	stmt, err := mysql.Client.Prepare(queryDeleteOrder)
	if err != nil {
		logger.Error("error when trying to prepare delete order statement", err)
		return responses.NewInternalServerError("Internal Server Error", "Please try again later...")
	}
	defer stmt.Close()

	_, delErr := stmt.Exec(order.UserId, order.OrderID)
	if delErr != nil {
		logger.Error("error when trying to delete order", delErr)
		return responses.NewInternalServerError("Internal Server Error", "Please try again later...")
	}

	return nil
}