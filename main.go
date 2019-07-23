package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

	if err != nil {
		log.Fatal(err)
	}

	gz, err := gzip.NewReader(file)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	defer gz.Close()

	fullStr, err := ioutil.ReadAll(gz)

	if err != nil {
		log.Fatal(err)
	}

	return fullStr
}

func main() {

	path := "data/"
	filename := "GameData.json.gz"

	dataBytes := GetGzText(path + filename)

	var jsonMap map[string]interface{}
	json.Unmarshal(dataBytes, &jsonMap)

	intCount := 0
	stringCount := 0
	nullCount := 0

	for element, value := range jsonMap {
		switch value.(type) {
		case float64:
			intCount++
			fmt.Println(element, "is Int")
		case string:
			stringCount++
			fmt.Println(element, "is String")
		case []interface{}:
			nullCount++
			fmt.Println(element, "is Array")
		default:
			nullCount++
			fmt.Println(element, "is Object I think.....")
		}
	}

	fmt.Println("Strings values ", stringCount)

	fmt.Println("Numeric values ", intCount)

	fmt.Println("Null value ", nullCount)
}
