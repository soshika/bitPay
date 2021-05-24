package app

import (
	"bitPay/controllers/binance"
	"bitPay/controllers/ping"
)

func urlPatterns() {
	router.GET("/ping", ping.Ping)

	router.POST("/balance", binance.GetBalance)
	router.POST("/getOrder", binance.GetOrder)
	router.POST("/createOrder", binance.CreateOrder)
	router.POST("/cancelOrder", binance.CancelOrder)
}
