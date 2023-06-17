package server

import (
	"game_with_Nikita/internal/localSite/adress"
	"github.com/labstack/echo/v4"
)

func Start() {
	myEcho := echo.New()

	/*echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())*/

	myEcho.GET("", adress.Handler)

	myEcho.Logger.Fatal(myEcho.Start(":8080"))
}
