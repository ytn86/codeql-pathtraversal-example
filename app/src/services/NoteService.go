package services

import (
	"io"
	"log"
	"os"
	"path"

	"github.com/google/uuid"
)

func ListNotes() ([]string, error) {

	notesPath := "notes/"
	files, err := os.ReadDir(notesPath)
	if err != nil {
		return nil, err
	}

	var noteList []string
	for _, file := range files {
		noteList = append(noteList, file.Name())
	}

	return noteList, nil
}

func GetNote(noteId string) ([]byte, error) {

	notePath := path.Join("notes/", noteId)
	log.Println(notePath)

	fp, err := os.Open(notePath)
	if err != nil {
		return nil, err
	}

	defer fp.Close()

	content, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func CreateNote(content string) (string, error) {

	noteId := uuid.New().String()
	notePath := path.Join("notes/", noteId)
	fp, err := os.Create(notePath)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	_, err = fp.WriteString(content)
	if err != nil {
		return "", err
	}

	return noteId, nil
}
