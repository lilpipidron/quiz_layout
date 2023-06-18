package adress

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"html/template"
	"net/http"
)

func Handler(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "Ты пидорас")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func SecondStage(ctx echo.Context) error {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Second quest</title>
		</head>
		<body>
			<img src="https://img.freepik.com/premium-vector/evil-pile-of-poo-brown-poop-with-grin-and-red-eyes-isolated-on-white-background_53562-14551.jpg" alt="My Image">
		</body>
		</html>
	`

	tmpl := template.Must(template.New("index").Parse(html))
	return tmpl.Execute(ctx.Response().Writer, nil)
}
