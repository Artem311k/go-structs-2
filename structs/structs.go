package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"kuzmin.com/structs2/fileInteractions"
	"kuzmin.com/structs2/utils"
)

type Note struct {
	Id        string
	Title     string
	Content   string
	CreatedAt string
}

func NewNoteFromInput() Note {
	return Note{
		utils.GenerateUUID(),
		utils.GetTitleFromInput(),
		utils.GetContentFromInpit(),
		utils.GetCreatedAt(),
	}
}

func (note *Note) ToString() string {
	return fmt.Sprintf("[%s, %s, %s, %s]", note.Id, note.Title, note.Content, note.CreatedAt)
}

func (note *Note) PrintNote() {
	fmt.Println(note.ToString())
}

func (note *Note) SaveNote() {
	fileInteractions.CreateDirectory()
	content, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(fileInteractions.GetNoteNameNumberWithPath(), content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
