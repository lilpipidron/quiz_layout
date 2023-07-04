package address

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"html/template"
	"net/http"
)

func Handler(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "Все работает")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func SecondStage(ctx echo.Context) error {
	//todo
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Second quest</title>
		</head>
		<body>

			<img src="суда путь на картинку" alt="если нет картинки">
		</body>
		</html>
	`

	tmpl := template.Must(template.New("index").Parse(html))
	return tmpl.Execute(ctx.Response().Writer, nil)
}
