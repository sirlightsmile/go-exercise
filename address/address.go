package address

//	- NewAddress(subDistrict, district, province, zipcode) - return new Address instance
//	- GetProvinces() - return list of all provinces
//  - GetDistrictsByProvince(province) - return list of all districts in a specified province
//  - GetZipcodesByDistrict(district) - return list of all zipcodes in a specified district
//
//  - create Address method:
//	- Validate() - return true if the address is valid

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Address struct {
	SubDistrict string
	District    string
	Province    string
	Zipcodes    string
}

func NewAddress(subDistrict string, district string, province string, zipcode string) {
}

func GetProvinces() {
	database, err := sql.Open("sqlite3", "data/th_address.db")
	checkErr(err)
	defer database.Close()

	rows, err := database.Query("SELECT PROVINCE_NAME_ENG FROM provinces")
	checkErr(err)

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		checkErr(err)
		fmt.Println(name)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
