package main

import (
	"kuzmin.com/structs2/fileInteractions"
	"kuzmin.com/structs2/structs"
)

func main() {

	newNote := structs.NewNoteFromInput()

	fileInteractions.Save(newNote)

	todo := structs.NewToDoFromInput()

	fileInteractions.Save(todo)
}
