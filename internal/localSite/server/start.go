package server

import (
	"game_with_Nikita/internal/localSite/adress"
	"github.com/labstack/echo/v4"
)

func Start(path string, stage int) {
	myEcho := echo.New()

	/*echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())*/

	myEcho.GET("", adress.Handler)

	pathFunc := adress.MakeFunc(123)

	myEcho.GET(path, pathFunc)

	myEcho.Logger.Fatal(myEcho.Start(":8080"))
}
