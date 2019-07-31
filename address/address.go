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

	_ "github.com/mattn/go-sqlite3"
)

func NewAddress(subDistrict string, district string, province string, zipcode string) {
}

func GetProvinces() []Province {
	database, err := sql.Open("sqlite3", "data/th_address.db")
	checkErr(err)
	defer database.Close()

	rows, err := database.Query("SELECT * FROM provinces")
	checkErr(err)

	var provinces []Province
	for rows.Next() {
		var province Province
		err = rows.Scan(&province.ID, &province.Code, &province.Name, &province.NameEng, &province.GeoID)
		checkErr(err)
		provinces = append(provinces, province)
	}
	rows.Close()

	return provinces
}

func GetDistrictsByProvince(province string) Amphur {

	database, err := sql.Open("sqlite3", "data/th_address.db")
	checkErr(err)
	defer database.Close()

	rows, err := database.Query("SELECT PROVINCE_ID FROM provinces WHERE PROVINCE_NAME IS ", province)
	checkErr(err)

	for rows.Next() {
	}

	var district Amphur
	return district
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
