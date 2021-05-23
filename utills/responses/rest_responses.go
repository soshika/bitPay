package responses

import (
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Header 	interface{}		`json:"header"`
	Body	interface{}		`json:"body"`
}


func (response *Response) Status() int {
	header 		:= response.Header.(map[string]interface{})
	statusCode 	:= header["status"].(int)

	return statusCode
}

func (response *Response) Message() string {
	header 	:= response.Header.(map[string]interface{})
	message	:= header["message"].(string)

	return message
}

func (response *Response) SetRequestId(requestId string) {
	header		:= response.Header.(map[string]interface{})
	header["requestId"] = requestId

	response.Header = header
}

func GetNow() string{
	return strings.Replace(strings.Replace(time.Now().UTC().Format("2006-01-02T15:04Z"), "T", " ", -1), "Z", "", -1)
}

func NewRequestSuccessfullyCreated(message string, detail string, body interface{}) *Response {

	header := make(map[string]interface{})

	header["status"]		= http.StatusCreated
	header["responseTime"]	= GetNow()
	header["message"]		= message
	header["detail"]		= detail

	return &Response{
		Header	: header,
		Body	: body,
	}
}

func NewRequestSuccessOk(message string, detail string, body interface{}) *Response {

	header := make(map[string]interface{})

	header["status"]		= http.StatusOK
	header["responseTime"]	= GetNow()
	header["message"]		= message
	header["detail"]		= detail

	return &Response{
		Header 	: header,
		Body	: body,
	}
}

func NewBadRequestError(message string, detail string, code int) *Response {

	header := make(map[string]interface{})

	header["status"]		= code
	header["responseTime"]	= GetNow()
	header["message"]		= message
	header["detail"]		= detail

	return &Response{
		Header	: header,
		Body	: nil,
	}
}

func NewInternalServerError(message string, detail string) *Response {

	header := make(map[string]interface{})

	header["status"]		= http.StatusInternalServerError
	header["responseTime"]	= GetNow()
	header["message"]		= message
	header["detail"]		= detail

	return &Response{
		Header	: header,
		Body	: nil,
	}
}

