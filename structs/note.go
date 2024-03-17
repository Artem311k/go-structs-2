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
	Id        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
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
	fileInteractions.CreateDirectory("note")
	content, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(fileInteractions.GetNameNumberWithPath("note"), content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
