package structs

import (
	"fmt"

	"kuzmin.com/structs2/utils"
)

type Note struct {
	UUUIDElement
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewNoteFromInput() Note {
	return Note{
		UUUIDElement: GenerateUUIDElement(),
		Title:        utils.GetTitleFromInput(),
		Content:      utils.GetContentFromInpit(),
	}
}

func (note *Note) ToString() string {
	return fmt.Sprintf("[%s, %s, %s, %s]", note.Id, note.Title, note.Content, note.CreatedAt)
}

func (note *Note) PrintInfo() {
	fmt.Println(note.ToString())
}
