package main

import (
	"bitPay/app"
	"github.com/gin-gonic/gin"
)

func main()  {
	gin.SetMode(gin.ReleaseMode)
	app.StartApplication()
}
