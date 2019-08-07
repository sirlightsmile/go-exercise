package main

//
// HTTP server
//
//
// Extend your exercise_02 and create a HTTP server that will listen on port 4000 and will implement these APIs:
//
//
// PREREQUISITES:
// - simplify your queries using joins and inner selects
// - implement repository pattern
//
//
// - Use HTTP GET method
// - Use go mod for dependencies (https://blog.golang.org/using-go-modules)
//
//
// 1) /api/validate_address
// 		with JSON payload in a request body, example:
//		{ "subDistrict": "Chan Kasem", "district": "Chatuchak", "province: "Bangkok", "zipcode": 10900 }
//
//  	if the address is valid, response will return HTTP status code 200 with response body { "result": "Address is valid." }
//		if the address is not valid, response will return HTTP status code 422 with response body { "error": "Address is not valid." }
//
//
// 2) /api/get_provinces
//		empty request body
//
//		will return HTTP status code 200 with response body { "result": [ "Bangkok", ... ] }
//
//
// 3) /api/get_district_by_province
// 		with JSON payload in a request body, example:
//		{ "province": "Chiang Mai" }
//
//		if province is valid, response will return HTTP status code 200 with response body { "result": [ ... districts ... ] }
//		if province is not valid, response will return status code 422 with response body { "error": "Province is not valid." }
//
//
// 4) /api/get_zipcodes_by_district
// 		with JSON payload in a request body, example:
//		{ "district": "Chatuchak" }
//
//		if district is valid, response will return HTTP status code 200 with response body { "result": [ 10000, 10100, ... ] }
//		if district is not valid, response will return status code 422 with response body { "error": "District is not valid." }
//
//
// 5) for any undefined url route return error HTTP status code 404 with response body { "error: "The requested URL was not found on the server." }
//  	examples of invalid urls:
//			/invalid_url
//			/api/invalid_url
//			/api
//			/
//
//
// 6) write test on all new APIs
//

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"smile/address"
)

const databasePath = "./data/th_address.db"

var db *address.SqlDB

func main() {

	absPath, err := filepath.Abs(databasePath)
	checkErr(err)
	db, err = address.ConnectSqlDB(absPath)
	checkErr(err)

	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/newAddress", NewAddress)
	http.HandleFunc("/validate", Validate)
	http.HandleFunc("/getProvinces", GetProvincesRequest)
	http.HandleFunc("/getDistrictByProvince", GetDistrictByProvince)
	http.HandleFunc("/getZipcodesByDistrict", GetZipcodesByDistrict)
	http.ListenAndServe(":4000", nil)
}

func GetProvincesRequest(w http.ResponseWriter, r *http.Request) {
	provinces, err := address.GetProvinces(db)
	checkErr((err))
	json.NewEncoder(w).Encode(provinces)
}

func GetDistrictByProvince(w http.ResponseWriter, r *http.Request) {
	type Task struct {
		Province string
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	districts, err := address.GetDistrictsByProvince(db, task.Province)
	checkErr((err))

	json.NewEncoder(w).Encode(districts)
}

func GetZipcodesByDistrict(w http.ResponseWriter, r *http.Request) {
	type Task struct {
		District string
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	zipcodes, err := address.GetZipcodesByDistrict(db, task.District)
	checkErr((err))

	json.NewEncoder(w).Encode(zipcodes)
}

func NewAddress(w http.ResponseWriter, r *http.Request) {
	var task struct {
		Province    string
		District    string
		SubDistrict string
		ZipCode     string
	}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	newAddress := address.NewAddress(db, task.SubDistrict, task.District, task.Province, task.ZipCode)

	json.NewEncoder(w).Encode(newAddress)
}

func Validate(w http.ResponseWriter, r *http.Request) {

	var task address.Address
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	valid := address.Validate(task)

	json.NewEncoder(w).Encode(valid)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
