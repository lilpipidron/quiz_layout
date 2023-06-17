package server

import (
	"game_with_Nikita/internal/localSite/adress"
	"github.com/labstack/echo/v4"
)

func Start() {
	echo := echo.New()

	/*echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())*/

	echo.GET("", adress.Handler)

	echo.Logger.Fatal(echo.Start(":8080"))
}
