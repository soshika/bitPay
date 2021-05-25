package binance

import (
	"bitPay/domains/binance"
	"bitPay/domains/users"
	"bitPay/logger"
	"bitPay/services"
	"bitPay/utills/parser"
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

func OpenOrders(c *gin.Context) {
	logger.Info("Enter to OpenOrders controller successfully")

	testnet 	:= c.GetHeader("testnet")
	apiKey 		:= c.GetHeader("api_key")
	secretKey 	:= c.GetHeader("secret_key")

	useTestnet, _ := strconv.ParseBool(testnet)

	client := users.User{
		APIKey: apiKey,
		SecretKey: secretKey,
	}

	openOrders := binance.OpenOrderRequest{}
	if err := c.ShouldBindJSON(&openOrders); err != nil {
		logger.Error("error when trying to bind json", err)
		restError := responses.NewBadRequestError("Invalid json structure","please make sure to send correct json body", http.StatusBadRequest)
		c.JSON(restError.Status(), restError)
		return
	}

	orders, serviceErr := services.BinancesService.OpenOrders(client, openOrders, useTestnet)
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), serviceErr)
		return
	}

	ok := responses.NewRequestSuccessOk("Open Orders sent Successfully.", "", orders)
	c.JSON(http.StatusOK,  ok)

	logger.Info("Close from OpenOrders controller successfully")
}

func ListOrders(c *gin.Context) {
	logger.Info("Enter to ListOrders controller successfully")

	testnet 	:= c.GetHeader("testnet")
	apiKey 		:= c.GetHeader("api_key")
	secretKey 	:= c.GetHeader("secret_key")

	useTestnet, _ := strconv.ParseBool(testnet)

	client := users.User{
		APIKey: apiKey,
		SecretKey: secretKey,
	}

	req := make(map[string]interface{})
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("error when trying to bind json", err)
		restError := responses.NewBadRequestError("Invalid json structure","please make sure to send correct json body", http.StatusBadRequest)
		c.JSON(restError.Status(), restError)
		return
	}

	symbol, parseErr := parser.ListOrderParser(req)
	if parseErr != nil {
		c.JSON(parseErr.Status(), parseErr)
		return
	}

	orders, serviceErr := services.BinancesService.ListOrders(client, useTestnet, *symbol)
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), serviceErr)
		return
	}

	ok := responses.NewRequestSuccessOk("Orders sent Successfully.", "", orders)
	c.JSON(http.StatusOK,  ok)

	logger.Info("Close from ListOrders controller successfully")
}