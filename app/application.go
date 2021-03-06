package app

import (
	"bitPay/logger"
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default()
)

func StartApplication()  {
	urlPatterns()
	logger.Info("about to start the application v 1.0.0 !")

	router.Run(":9092")
}
