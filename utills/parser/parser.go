package parser

import (
	"bitPay/utills/responses"
	"net/http"
)

func ListOrderParser(req map[string]interface{}) (*string, *responses.Response){
	symbol, found := req["symbol"].(string)
	if found == false {
		return nil, responses.NewBadRequestError("Bad Request", "Could not find symbol key in request", http.StatusBadRequest)
	}
	return &symbol, nil
}
