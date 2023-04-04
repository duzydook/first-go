package api

import (
	controller "demo/api/controllers"

	"github.com/labstack/echo/v4"
)

func DataAPI(g *echo.Group) {

	// crypto api
	g.GET("/cryptos", controller.RetrieveCryptos)

	// weather api
	g.GET("/weathers", controller.RetrieveWeathers)
	g.GET("/weathers/search", controller.SearchWeathers)

}
