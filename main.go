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

	data, err := ioutil.ReadAll(gz)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetJsonMap(bytes []byte) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal(bytes, &result)
	return result
}

func main() {

	path := "data/"
	filename := "GameData.json.gz"

	dataBytes := GetGzText(path + filename)

	var jsonMap map[string]interface{}
	jsonMap = GetJsonMap(dataBytes)

	fmt.Println(jsonMap["achievement"])
	/*
		scanner := bufio.NewScanner(gz)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			// Get Bytes and display the byte.
			b := scanner.Bytes()
			fmt.Printf("%v = %v = %v\n", b, b[0], string(b))
		}


			var result map[string]interface{}
			data.Unmarshal(scanner.Bytes(), &result)

			fmt.Println(result["achievement"])*/
}
