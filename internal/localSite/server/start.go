package server

import (
	"game_with_Nikita/internal/localSite/adress"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(path string, stage int) {
	myEcho := echo.New()

	myEcho.Use(middleware.Logger())
	myEcho.Use(middleware.Recover())

	myEcho.GET("", adress.Handler)
	myEcho.GET("/second", adress.SecondStage)

	myEcho.Logger.Fatal(myEcho.Start(":8080"))
}
