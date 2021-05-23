package app

import (
	"bitPay/controllers/ping"
	"bitPay/controllers/user"
)

func urlPatterns() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/balance", user.GetBalance)
}
