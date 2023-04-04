package main

import (
	"demo/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// groups
	dataGroup := e.Group("/data")

	api.DataAPI(dataGroup)
	api.Websocket(e)
	api.Line(e)

	e.Logger.Fatal(e.Start(":8888"))

}
