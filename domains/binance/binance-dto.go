package binance

type CreateOrderRequest struct {
	UserId 		int64			`json:"user_id"`
	Symbol 		string 			`json:"symbol"`
	Quantity	string 			`json:"quantity"`
	Price		string 			`json:"price"`
}

type GetOrderRequest struct {
	UserId 		int64		`json:"user_id"`
	OrderId		int64 		`json:"order_id"`
	Symbol		string 		`json:"symbol"`
}

type Order struct {
	Id 							int64		`json:"id"`
	UserId 						int64 		`json:"user_id"`
	Symbol     					string 		`json:"symbol"`
	OrderID     				int64  		`json:"order_id"`
	ClientOrderID           	string 		`json:"client_order_id"`
	TransactTime             	int64  		`json:"transact_time"`
	Price                    	string 		`json:"price"`
}

type OpenOrderRequest struct {
	UserId 		int64		`json:"user_id"`
	Symbol 		string 		`json:"symbol"`
}