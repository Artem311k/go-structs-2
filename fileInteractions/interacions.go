package fileInteractions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"kuzmin.com/structs2/utils"
)

func CreateDirectory(pattern string) {

	resolveDir(&pattern)

	cwd, _ := os.Getwd()

	if _, err := os.Stat(cwd + "/" + pattern); os.IsNotExist(err) {
		err := os.Mkdir(pattern, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		return
	}
}

func GetNameNumberWithPath(pattern string) string {

	counter := 1

	fileName := fmt.Sprintf(pattern+"%d.json", counter)

	cwd, _ := os.Getwd()

	filePath := buildFilePath(cwd, fileName, pattern)

	_, err := os.Stat(filePath)

	for err == nil {
		counter++
		fileName = fmt.Sprintf(pattern+"%d.json", counter)
		filePath = buildFilePath(cwd, fileName, pattern)
		_, err = os.Stat(filePath)
	}

	return filePath
}

func buildFilePath(cwd, fileName, pattern string) string {
	resolveDir(&pattern)
	return fmt.Sprintf("%s/%s/%s", cwd, pattern, fileName)
}

func resolveDir(pattern *string) {
	if *pattern == "Note" {
		*pattern = utils.NOTES_DIR_NAME
	} else if *pattern == "ToDo" {
		*pattern = utils.TODO_DIR_NAME
	} else {
		*pattern = utils.UNKNOWN_DIR_NAME
	}
}

func Save(data any) {
	typeNme := utils.GetType(data)
	CreateDirectory(typeNme)
	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(GetNameNumberWithPath(typeNme), content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
