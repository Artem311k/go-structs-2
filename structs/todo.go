package structs

import (
	"fmt"

	"kuzmin.com/structs2/utils"
)

type ToDo struct {
	UUUIDElement
	Content string `json:"content"`
}

func NewToDoFromInput() ToDo {
	return ToDo{
		UUUIDElement: GenerateUUIDElement(),
		Content:      utils.GetContentFromInpit(),
	}
}

func (todo *ToDo) ToString() string {
	return fmt.Sprintf("[%s, %s, %s]", todo.Id, todo.Content, todo.CreatedAt)
}

func (todo *ToDo) PrintInfo() {
	fmt.Println(todo.ToString())
}
