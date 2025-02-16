package routes

import (
	"github.com/labstack/echo"
	"lfiapp/api"
)

func Router(e *echo.Echo) {
	e.GET("/api/notes", api.GetNote)
	e.POST("/api/notes", api.CreateNote)
	e.GET("/api/notes/list", api.ListNotes)
}
