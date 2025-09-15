package routes

import (
	"github.com/labstack/echo"
	"pathtravapp/api"
)

func Router(e *echo.Echo) {
	e.GET("/api/notes", api.GetNote)
	e.POST("/api/notes", api.CreateNote)
	e.GET("/api/notes/list", api.ListNotes)
}
