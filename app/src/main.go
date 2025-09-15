package main

import (
	"github.com/labstack/echo"
	"pathtravapp/routes"
)

func main() {
	e := echo.New()
	routes.Router(e)

	e.Start(":5000")
}
