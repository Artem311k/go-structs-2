package fileInteractions

import (
	"fmt"
	"os"

	"kuzmin.com/structs2/utils"
)

func CreateDirectory() {
	cwd, _ := os.Getwd()

	if _, err := os.Stat(cwd + "/" + utils.NOTES_DIR_NAME); os.IsNotExist(err) {
		err := os.Mkdir(utils.NOTES_DIR_NAME, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		return
	}
}

func GetNoteNameNumberWithPath() string {

	counter := 1

	fileName := fmt.Sprintf("note%d.json", counter)

	cwd, _ := os.Getwd()

	filePath := buildFilePath(cwd, fileName)

	_, err := os.Stat(filePath)

	for err == nil {
		counter++
		fileName = fmt.Sprintf("note%d.json", counter)
		filePath = buildFilePath(cwd, fileName)
		_, err = os.Stat(filePath)
	}

	return filePath
}

func buildFilePath(cwd string, fileName string) string {
	return fmt.Sprintf("%s/%s/%s", cwd, utils.NOTES_DIR_NAME, fileName)
}
