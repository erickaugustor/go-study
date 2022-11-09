package main

import (
	"echo-api/src/router"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.Build(e)

	e.Logger.Fatal(e.Start(":1323"))
}
