package web

import (
	"aviasales/cache"
	"aviasales/config"
	"aviasales/modules/aviasales"
	"aviasales/web/handlers/places"
	projectMiddleware "aviasales/web/middlewares"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

var (
	e    *echo.Echo
	conf *config.Config
)

func init() {
	e = echo.New()
}

func StartUp(filepath string) {
	conf = config.Manager(filepath)

	e.Debug = conf.Debug
	client := cache.Manager(&conf.Redis)
	aviasalesService := aviasales.Manager(&conf.Avisales)
	c := projectMiddleware.NewAviasalesContext(conf, client, aviasalesService)
	e.Use(c.Process)
	api := e.Group("/api")
	{
		v1Api := api.Group("/v1")
		{
			placesV1Api := v1Api.Group("/places")
			{
				placesV1Api.GET("", places.SearchPlaces, middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
					if c.Response().Status == http.StatusOK && c.Request().Method == http.MethodGet {
						cc := c.(*projectMiddleware.AviaSalesContext)
						_, err := cc.Redis.Set(c.Request().RequestURI, resBody, conf.Redis.Expiration).Result()
						if err != nil {
							cc.Logger().Warn(err)
						}
					}
				}))
				placesV1Api.GET("/locales", places.AvailableLocales)
				placesV1Api.GET("/types", places.AvailableSearchTypes)
			}
		}
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Server.Port)))
}
