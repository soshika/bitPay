package redis

import (
	"bitPay/logger"
	"bitPay/utills/responses"
	"github.com/gomodule/redigo/redis"
)

func ConvertMinutesToSeconds(minutes int) int {
	return minutes * 60
}

func SaveWithTimeOut(key string, data string, timeout int, c redis.Conn) *responses.Response {
	_, err := c.Do("SETEX", key, timeout, data)
	if err != nil {
		logger.Error("error when trying to save into redis", err)
		return responses.NewInternalServerError("Could not save into redis", "")
	}

	return nil
}

func SaveWithOutTimeOut(key string, data string, c redis.Conn) *responses.Response {
	_, err := c.Do("SET", key, data)
	if err != nil {
		logger.Error("error when trying to save into redis", err)
		return responses.NewInternalServerError("Could not save into redis", "")
	}

	return nil
}

func Get(key string, c redis.Conn) (*string, *responses.Response){
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		switch err {
		case redis.ErrNil:
			logger.Error("error when trying to get data from redis, key does not exist", err)
		case redis.ErrPoolExhausted:
			logger.Error("error when trying to get data from redis, Expire Error", err)
		}
		return nil, responses.NewInternalServerError("Internal Server Error", "Please Try again later...")
	}
	return &s, nil

}

