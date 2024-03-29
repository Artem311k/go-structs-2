package utils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/google/uuid"
)

func GetStringFromInput(displayText string) string {
	fmt.Print(displayText)
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}
	return text
}

func GetTitleFromInput() string {
	title := GetStringFromInput("Enter title: ")
	if !ValidateTitle(title) {
		fmt.Println("Too large title, try again...")
		title = GetStringFromInput("Enter title: ")
	}
	return title
}

func GetContentFromInpit() string {
	content := GetStringFromInput("Enter content: ")
	if !ValidateContent(content) {
		fmt.Println("Too large content title, try again...")
		content = GetStringFromInput("Enter content: ")
	}
	return content
}

func GenerateUUID() string {
	return uuid.New().String()
}

func GetCreatedAt() string {
	return time.Now().Format("02.01.2006 15:04")
}

func GetType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
