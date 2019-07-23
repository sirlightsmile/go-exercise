package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
)

//
// GameData analyzer
//
// create program that open game data file 'data/GameData.json.gz', parse and analyze the content
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

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	path := "data/"
	filename := "GameData.json.gz"
	file, err := os.Open(path + filename)

	if err != nil {
		log.Fatal(err)
	}

	gz, err := gzip.NewReader(file)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	defer gz.Close()

	scanner := bufio.NewScanner(gz)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
