package middlewares

import (
	"aviasales/config"
	"aviasales/modules/aviasales"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

type AviaSalesContext struct {
	echo.Context
	Configuration    *config.Config
	Redis            *redis.Client
	AviasalesService *aviasales.Service
}

func NewAviasalesContext(conf *config.Config, client *redis.Client, service *aviasales.Service) *AviaSalesContext {
	return &AviaSalesContext{
		Configuration:    conf,
		Redis:            client,
		AviasalesService: service,
	}
}

func (cc *AviaSalesContext) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc.Context = c
		return next(cc)
	}
}

func (cc *AviaSalesContext) GetCachedResponse(key string) (interface{}, error) {
	var (
		err      error
		data     string
		response interface{}
	)
	data, err = cc.Redis.Get(key).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
