package server

import (
	"game_with_Nikita/internal/localsite/address"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(path string, stage int) {
	myEcho := echo.New()

	myEcho.Use(middleware.Logger())
	myEcho.Use(middleware.Recover())

	myEcho.GET("", address.Handler)
	myEcho.GET("/second", address.SecondStage)

	myEcho.Logger.Fatal(myEcho.Start(":8080"))
}
