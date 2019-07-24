package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

//
// Gamedata analyzer
//
// create program that open game data file 'data/Gamedata.data.gz', parse and analyze the content
// the program should print:
//
// 1) compressed file size in bytes
// 2) uncompressed size in bytes
// 3) number of game data tables
// 4) number of string values
// 5) number of numerical values (Note: numerical value can be 1.0, 5 or "10" or '5.45')
// 6) number of null values
//
//
// example output:
//
//
// Compressed size: 1234
// Uncompressed size: 1234
// Tables: 1234
// Strings values: 1234
// Numeric values: 1234
// Null values: 1234
//
const gameDataPath = "data/GameData.json.gz"

func main() {

	dataBytes := GetGzText(gameDataPath)

	var jsonData interface{}
	err := json.Unmarshal(dataBytes, &jsonData)

	jsonMap, ok := jsonData.(map[string]interface{})

	if !ok {
		fmt.Println("cannot convert the JSON objects")
		os.Exit(1)
	}

	LogErrorHandler(err)

	intCount := 0
	stringCount := 0
	nilCount := 0

	for element := range jsonMap {
		CountType(jsonMap[element], &intCount, &stringCount, &nilCount)
	}

	fmt.Println("Table count ", len(jsonMap))

	fmt.Println("Strings values ", stringCount)

	fmt.Println("Numeric values ", intCount)

	fmt.Println("Null value ", nilCount)
}

func GetGzText(filePath string) []byte {

	file, err := os.Open(filePath)

	LogErrorHandler(err)

	fileStat, err := file.Stat()

	LogErrorHandler(err)

	fmt.Println("Compressed File size : ", fileStat.Size())

	gz, err := gzip.NewReader(file)

	LogErrorHandler(err)

	defer file.Close()
	defer gz.Close()

	fullStr, err := ioutil.ReadAll(gz)

	LogErrorHandler(err)

	fmt.Println("Decompressed File size : ", len(fullStr))

	return fullStr
}

func LogErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CountType(obj interface{}, intCount *int, stringCount *int, nilCount *int) {

	switch obj.(type) {
	case float64, int:
		*intCount++

	case string:
		*stringCount++

	case nil:
		*nilCount++

	case []interface{}:
		objSlice := obj.([]interface{})
		for i := range objSlice {
			CountType(objSlice[i], &*intCount, &*stringCount, &*nilCount)
		}

	case map[string]interface{}:

		objMap := obj.(map[string]interface{})
		for element := range objMap {
			CountType(objMap[element], &*intCount, &*stringCount, &*nilCount)
		}

	default:
		fmt.Println("I don't know how to handler it.", reflect.TypeOf(obj))
	}
}
