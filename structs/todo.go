package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"kuzmin.com/structs2/fileInteractions"
	"kuzmin.com/structs2/utils"
)

type ToDo struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

func NewToDoFromInput() ToDo {
	return ToDo{
		utils.GenerateUUID(),
		utils.GetContentFromInpit(),
		utils.GetCreatedAt(),
	}
}

func (todo *ToDo) ToString() string {
	return fmt.Sprintf("[%s, %s, %s", todo.Id, todo.Content, todo.CreatedAt)
}

func (todo *ToDo) PrintToDo() {
	fmt.Println(todo.ToString())
}

func (todo *ToDo) SaveToDo() {
	fileInteractions.CreateDirectory("todo")
	content, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(fileInteractions.GetNameNumberWithPath("todo"), content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
