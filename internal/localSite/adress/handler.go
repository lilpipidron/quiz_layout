package adress

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

func Handler(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "Ты пидорас")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func MakeFunc(stage int) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		err := ctx.File("admin123-log.txt")
		if err != nil {
			log.Fatal(err)
		}
		err = ctx.String(http.StatusOK, strconv.Itoa(stage))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
}
