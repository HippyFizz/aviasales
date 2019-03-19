package places

import (
	"aviasales/modules/aviasales"
	"aviasales/web/middlewares"
	"github.com/labstack/echo"
	"net/http"
)

func SearchPlaces(c echo.Context) error {
	var response []*aviasales.WidgetFormat
	cc := c.(*middlewares.AviaSalesContext)
	cache, err := cc.GetCachedResponse(c.Request().RequestURI)

	if err == nil {
		c.Logger().Info("Request was loaded from cache")
		return c.JSON(http.StatusOK, cache)
	}

	values := c.QueryParams()
	response, err = cc.AviasalesService.RequestPlaces(values, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

func AvailableLocales(c echo.Context) error {
	cc := c.(*middlewares.AviaSalesContext)
	return c.JSON(http.StatusOK, cc.AviasalesService.Locales)
}

func AvailableSearchTypes(c echo.Context) error {
	cc := c.(*middlewares.AviaSalesContext)
	return c.JSON(http.StatusOK, cc.AviasalesService.Types)
}
