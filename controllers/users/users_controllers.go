package users

import (
	"bitPay/domains/users"
	"bitPay/logger"
	"bitPay/services"
	"bitPay/utills/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserRegistration(c *gin.Context) {
	logger.Info("Enter to UserRegistration controller successfully")

	testnet 	:= c.GetHeader("testnet")
	apiKey 		:= c.GetHeader("api_key")
	secretKey 	:= c.GetHeader("secret_key")

	useTestnet, _ := strconv.ParseBool(testnet)

	client := users.User{
		IsTest: useTestnet,
		APIKey: apiKey,
		SecretKey: secretKey,
	}

	serviceErr := services.UsersService.Register(client)
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), serviceErr)
		return
	}

	c.JSON(http.StatusOK, responses.NewRequestSuccessOk("user registered successfully", "", nil))

	logger.Info("Close from UserRegistration controller successfully")
}