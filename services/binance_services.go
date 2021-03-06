package services

import (
	binance2 "bitPay/domains/binance"
	"bitPay/domains/users"
	"bitPay/logger"
	"bitPay/utills/responses"
	"context"
	"github.com/adshao/go-binance/v2"
	"net/http"
)

var (
	BinancesService		binanceServiceInterface = & binancesService{}
)

type binancesService struct {}

type binanceServiceInterface interface {
	GetBalance		(users.User, string, bool)							(*binance.Balance, *responses.Response)
	CreateOrder		(users.User, binance2.CreateOrderRequest, bool) 	(*binance2.Order, *responses.Response)
	GetOrder		(users.User, binance2.GetOrderRequest, bool)		(*binance.Order, *responses.Response)
	CancelOrder		(users.User, binance2.GetOrderRequest, bool) 		*responses.Response
	OpenOrders		(users.User, binance2.OpenOrderRequest, bool) 		([]binance2.Order, *responses.Response)
	ListOrders		(users.User, bool, string) 							([]*binance.Order, *responses.Response)

}

func (s *binancesService) CreateOrder(user users.User, createOrder binance2.CreateOrderRequest, useTestnet bool) (*binance2.Order, *responses.Response) {
	logger.Info("Enter to CreateOrder service successfully")
	if useTestnet {
		binance.UseTestnet = useTestnet
	}
	client := binance.NewClient(user.APIKey, user.SecretKey)

	orderResponse, err := client.NewCreateOrderService().Symbol(createOrder.Symbol).
		Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity(createOrder.Quantity).
		Price(createOrder.Price).Do(context.Background())
	if err != nil {
		return nil, responses.NewBadRequestError(err.Error(), "Could not create order, please try again later...", http.StatusBadRequest)
	}

	order := binance2.Order{
		UserId: user.Id,
		Symbol: orderResponse.Symbol,
		OrderID: orderResponse.OrderID,
		ClientOrderID: orderResponse.ClientOrderID,
		TransactTime: orderResponse.TransactTime,
		Price: orderResponse.Price,
	}

	_ = order.Save()

	logger.Info("Close from CreateOrder service successfully")
	return &order, nil
}

func (s *binancesService) GetOrder(user users.User, getOrder binance2.GetOrderRequest, useTestnet bool) (*binance.Order, *responses.Response) {
	logger.Info("Enter to GetOrder service successfully")

	if useTestnet {
		binance.UseTestnet = true
	}

	client := binance.NewClient(user.APIKey, user.SecretKey)

	order := binance2.Order{
		UserId: user.Id,
		OrderID: getOrder.OrderId,
	}

	getErr := order.Get()
	if getErr != nil {
		return nil, getErr
	}

	orderResponse, err := client.NewGetOrderService().Symbol(order.Symbol).
		OrderID(order.Id).Do(context.Background())
	if err != nil {
		return nil, responses.NewBadRequestError("Bad Request", err.Error(), http.StatusBadRequest)
	}

	logger.Info("Close from GetOrder service successfully")
	return orderResponse, nil
}

func (s *binancesService) CancelOrder(user users.User, getOrder binance2.GetOrderRequest, useTestnet bool) *responses.Response {
	logger.Info("Enter to CancelOrder service successfully")

	if useTestnet {
		binance.UseTestnet = true
	}

	client := binance.NewClient(user.APIKey, user.SecretKey)

	order := binance2.Order{
		UserId: user.Id,
		OrderID: getOrder.OrderId,
	}

	getErr := order.Get()
	if getErr != nil {
		return getErr
	}

	_, err := client.NewCancelOrderService().Symbol(order.Symbol).
		OrderID(order.OrderID).Do(context.Background())
	if err != nil {
		return responses.NewBadRequestError("Bad request", err.Error(), http.StatusBadRequest)
	}

	_  = order.Delete()

	logger.Info("Close from CancelOrder service successfully")
	return nil
}

func (s *binancesService) OpenOrders(user users.User, AssetOrders binance2.OpenOrderRequest, useTestnet bool) ([]binance2.Order, *responses.Response) {
	logger.Info("Enter to OpenOrders service successfully")

	if getErr := user.Get(); getErr != nil {
		return nil, getErr
	}

	order := binance2.Order{
		UserId: user.Id,
		Symbol: AssetOrders.Symbol,
	}

	orders, getErr := order.GetAll()
	if getErr != nil {
		return nil, getErr
	}

	logger.Info("Close from OpenOrders service successfully")
	return orders, nil
}

func (s *binancesService) ListOrders(user users.User, useTestnet bool, symbol string) ([]*binance.Order, *responses.Response) {
	logger.Info("Enter to ListOrders service successfully")

	if useTestnet {
		binance.UseTestnet = true
	}

	client := binance.NewClient(user.APIKey, user.SecretKey)

	orders, err := client.NewListOrdersService().Symbol(symbol).
		Do(context.Background())
	if err != nil {
		return nil, responses.NewBadRequestError("Bad request", err.Error(), http.StatusBadRequest)
	}

	logger.Info("Close from ListOrders service successfully")
	return orders, nil
}

func(s *binancesService) GetBalance(client users.User, currency string, useTestnet bool) (*binance.Balance, *responses.Response){
	logger.Info("Enter To GetBalance service successfully.")

	if useTestnet {
		binance.UseTestnet = true
	}

	binance := binance.NewClient(client.APIKey, client.SecretKey)
	res, err := binance.NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, responses.NewBadRequestError(err.Error(), "Bad request", http.StatusBadRequest)
	}

	for _, balance := range res.Balances {
		if balance.Asset == currency {
			return &balance, nil
		}
	}

	logger.Info("Close from GetBalance service successfully.")
	return nil, nil
}
