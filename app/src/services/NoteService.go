package services

import (
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"path"
)

func ListNotes() ([]string, error) {

	notes_path := "notes/"
	files, err := os.ReadDir(notes_path)
	if err != nil {
		return nil, err
	}

	var note_list []string
	for _, file := range files {
		note_list = append(note_list, file.Name())
	}

	return note_list, nil
}

func GetNote(noteId string) ([]byte, error) {

	note_path := path.Join("notes/", noteId)
	log.Println(note_path)

	fp, err := os.Open(note_path)
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
	note_path := path.Join("notes/", noteId)
	fp, err := os.Create(note_path)
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
