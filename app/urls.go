package app

import (
	"bitPay/controllers/binance"
	"bitPay/controllers/ping"
	"bitPay/controllers/users"
)

func urlPatterns() {
	router.GET("/ping", ping.Ping)

	router.POST("/balance", binance.GetBalance)
	router.POST("/getOrder", binance.GetOrder)
	router.POST("/createOrder", binance.CreateOrder)
	router.POST("/cancelOrder", binance.CancelOrder)
	router.POST("/openOrders", binance.OpenOrders)
	router.POST("/orders", binance.ListOrders)

	router.POST("user/register", users.UserRegistration)
}
