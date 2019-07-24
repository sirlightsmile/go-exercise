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

func main() {

	path := "data/"
	filename := "GameData.json.gz"

	dataBytes := GetGzText(path + filename)

	var jsonData interface{}
	err := json.Unmarshal(dataBytes, &jsonData)

	jsonMap, ok := jsonData.(map[string]interface{})

	if !ok {
		fmt.Println("cannot convert the JSON objects")
		os.Exit(1)
	}

	LogErrorHandler(err)

	intCount, stringCount, nilCount := CountMapType(jsonMap)

	fmt.Println("Table count ", len(jsonMap))

	fmt.Println("Strings values ", stringCount)

	fmt.Println("Numeric values ", intCount)

	fmt.Println("Null value ", nilCount)
}

func CountMapType(jsonMap map[string]interface{}) (int, int, int) {

	intCount := 0
	stringCount := 0
	nilCount := 0

	for element := range jsonMap {
		switch jsonMap[element].(type) {
		case float64:
			intCount++

		case string:
			stringCount++

		case nil:
			nilCount++

		case []interface{}:
			intPlus, stringPlus, nilPlus := CountInterfaceType(jsonMap[element].([]interface{}))
			intCount += intPlus
			stringCount += stringPlus
			nilCount += nilPlus

		case map[string]interface{}:

			intPlus, stringPlus, nilPlus := CountMapType(jsonMap[element].(map[string]interface{}))
			intCount += intPlus
			stringCount += stringPlus
			nilCount += nilPlus

		default:
			fmt.Println("I don't know how to handler it.", reflect.TypeOf(jsonMap[element]))
		}
	}

	return intCount, stringCount, nilCount
}

func CountInterfaceType(jsonMap []interface{}) (int, int, int) {

	intCount := 0
	stringCount := 0
	nilCount := 0

	for element := range jsonMap {
		switch jsonMap[element].(type) {
		case float64:
			intCount++

		case string:
			stringCount++

		case nil:
			nilCount++

		case []interface{}:
			intPlus, stringPlus, nilPlus := CountInterfaceType(jsonMap[element].([]interface{}))
			intCount += intPlus
			stringCount += stringPlus
			nilCount += nilPlus

		case map[string]interface{}:

			intPlus, stringPlus, nilPlus := CountMapType(jsonMap[element].(map[string]interface{}))
			intCount += intPlus
			stringCount += stringPlus
			nilCount += nilPlus

		default:
			fmt.Println("I don't know how to handler it.", reflect.TypeOf(jsonMap[element]))
		}
	}

	return intCount, stringCount, nilCount
}
