package binance

import (
	"bitPay/domains/binance"
	"bitPay/domains/users"
	"bitPay/logger"
	"bitPay/services"
	"bitPay/utills/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateOrder(c *gin.Context) {
	logger.Info("Enter to CreateOrder controller successfully")

	testnet 	:= c.GetHeader("testnet")
	apiKey 		:= c.GetHeader("api_key")
	secretKey 	:= c.GetHeader("secret_key")


	useTestnet, _ := strconv.ParseBool(testnet)

	client := users.User{
		APIKey: apiKey,
		SecretKey: secretKey,
	}

	createOrder := binance.CreateOrderRequest{}
	if err := c.ShouldBindJSON(&createOrder); err != nil {
		logger.Error("error when trying to bind json", err)
		restError := responses.NewBadRequestError("Invalid json structure","please make sure to send correct json body", http.StatusBadRequest)
		c.JSON(restError.Status(), restError)
		return
	}

	order, serviceErr := services.BinancesService.CreateOrder(client, createOrder, useTestnet)
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), serviceErr)
		return
	}

	ok := responses.NewRequestSuccessOk("Order Created Successfully.", "", order)
	c.JSON(http.StatusOK,  ok)

	logger.Info("Close from CreateOrder controller successfully")
}

func GetBalance(c *gin.Context) {
	logger.Info("Enter to CreateAgency controller successfully")

	testnet 	:= c.GetHeader("testnet")
	apiKey 		:= c.GetHeader("api_key")
	secretKey 	:= c.GetHeader("secret_key")

	useTestnet, _ := strconv.ParseBool(testnet)

	client := users.User{
		APIKey: apiKey,
		SecretKey: secretKey,
	}

	balance, serviceErr := services.BinancesService.GetBalance(client, "USDT", useTestnet)
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), serviceErr)
		return
	}

	ok := responses.NewRequestSuccessOk("balance sent successfully", "", balance)
	c.JSON(http.StatusOK,  ok)

	logger.Info("Close from CreateAgency controller successfully")
}

func GetOrder(c *gin.Context) {
	logger.Info("Enter into GetOrder controller successfully")

	testnet 	:= c.GetHeader("testnet")
	apiKey 		:= c.GetHeader("api_key")
	secretKey 	:= c.GetHeader("secret_key")

	useTestnet, _ := strconv.ParseBool(testnet)

	client := users.User{
		APIKey: apiKey,
		SecretKey: secretKey,
	}

	getOrder := binance.GetOrderRequest{}
	if err := c.ShouldBindJSON(&getOrder); err != nil {
		logger.Error("error when trying to bind json", err)
		restError := responses.NewBadRequestError("Invalid json structure","please make sure to send correct json body", http.StatusBadRequest)
		c.JSON(restError.Status(), restError)
		return
	}

	order, serviceErr := services.BinancesService.GetOrder(client, getOrder, useTestnet)
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), serviceErr)
		return
	}

	ok := responses.NewRequestSuccessOk("Order information sent Successfully.", "", order)
	c.JSON(http.StatusOK,  ok)

	logger.Info("Close from GetOrder controller successfully")
}

func CancelOrder(c *gin.Context) {
	logger.Info("Enter to CancelOrder controller successfully")

	testnet 	:= c.GetHeader("testnet")
	apiKey 		:= c.GetHeader("api_key")
	secretKey 	:= c.GetHeader("secret_key")

	useTestnet, _ := strconv.ParseBool(testnet)

	client := users.User{
		APIKey: apiKey,
		SecretKey: secretKey,
	}

	getOrder := binance.GetOrderRequest{}
	if err := c.ShouldBindJSON(&getOrder); err != nil {
		logger.Error("error when trying to bind json", err)
		restError := responses.NewBadRequestError("Invalid json structure","please make sure to send correct json body", http.StatusBadRequest)
		c.JSON(restError.Status(), restError)
		return
	}

	serviceErr := services.BinancesService.CancelOrder(client, getOrder, useTestnet)
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), serviceErr)
		return
	}

	ok := responses.NewRequestSuccessOk("Order canceled Successfully.", "", nil)
	c.JSON(http.StatusOK,  ok)


	logger.Info("Close from CancelOrder controller successfully")
}