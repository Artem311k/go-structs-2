package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func getConfigMap() map[string]interface{} {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Print("Can not find config file")
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	return result
}

func ReadStringConfig(configName string) string {

	return getConfigMap()[configName].(string)

}

func ReadIntConfig(configName string) int {

	return int(getConfigMap()[configName].(float64))

}
