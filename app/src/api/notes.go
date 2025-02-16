package api

import (
	"github.com/labstack/echo"
	"lfiapp/services"
	"net/http"
)

func GetNote(c echo.Context) error {

	noteId := c.QueryParam("id")
	content, err := services.GetNote(noteId)

	if err != nil {
		return c.String(http.StatusNotFound, "Note not found")
	}

	return c.String(http.StatusOK, string(content))
}

func CreateNote(c echo.Context) error {
	content := c.FormValue("content")
	noteId, err := services.CreateNote(content)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create note")
	}

	return c.String(http.StatusOK, noteId)
}

func ListNotes(c echo.Context) error {

	note_list, err := services.ListNotes()

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to list notes")
	}

	return c.JSON(http.StatusOK, note_list)
}
