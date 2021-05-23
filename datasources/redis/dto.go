package redis

import (
	"bitPay/logger"
	"bitPay/utills/responses"
	"github.com/gomodule/redigo/redis"
)

func Init() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			//c, err := redis.Dial("tcp", ":6379")
			//if err != nil {
			//	panic(err.Error())
			//}
			//return c, err
			//return redis.DialURL("redis://@E@8@3@3@x@s@b@f@@localhost:6379")
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", "@E@8@3@3@x@s@b@f@"); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func Ping(c redis.Conn) *responses.Response {
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		logger.Error("error when trying to send ping to redis", err)
		return responses.NewInternalServerError("Internal Server Error", "Please Try Again later...")
	}

	logger.Info("PING = " + s)

	return nil
}

