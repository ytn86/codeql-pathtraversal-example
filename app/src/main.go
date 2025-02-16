package main

import (
	"github.com/labstack/echo"
	"lfiapp/routes"
)

func main() {
	e := echo.New()
	routes.Router(e)

	e.Start(":5000")
}
