package main

//
// Address package
//
// Create 'address' package for lookup and validation of Thai address over a database
// this database is commonly used in e-commerce businesses in Thailand for address validation
//
// Requirements:
//
// 1) package name must be called 'address'
//
//
// 2) There is a database (SQLite) in folder 'data/th_address.db'
//	- access this database and use it for all queries
//
//
// 3) define structure Address in the package with these properties:
//		- subDistrict (districts table)
//		- district (amphures table)
//		- province (provinces table)
//		- zipcode (zipcodes table)
//
//
// 4) create these functions inside the package
//
//	- NewAddress(subDistrict, district, province, zipcode) - return new Address instance
//	- GetProvinces() - return list of all provinces
//  - GetDistrictsByProvince(province) - return list of all districts in a specified province
//  - GetZipcodesByDistrict(district) - return list of all zipcodes in a specified district
//
//  - create Address method:
//	- Validate() - return true if the address is valid
//
//
// 6) create test file in the same package ('address_test.go')
//  - write test cases for all methods above
//  - test coverage must be 100% (go test -cover)
//  - all tests must pass
//
//
// Additional notes: Use only English names, address validation is NOT case sensitive
//
//

import (
	"fmt"

	"./address"
)

func main() {
	provinces, err := address.GetProvinces()
	checkErr(err)
	for _, element := range provinces {
		fmt.Println(element.Name)
	}

	fmt.Println("===== Amphur =====")
	districts, err := address.GetDistrictsByProvince(provinces[0])
	checkErr(err)
	for _, element := range districts {
		fmt.Println(element.Name)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
