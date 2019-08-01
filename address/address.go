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

func GetProvinces() ([]Province, error) {
	database, err := sql.Open("sqlite3", "data/th_address.db")
	checkErr(err)
	defer database.Close()

	rows, err := database.Query("SELECT * FROM provinces")
	if err != nil {
		return nil, err
	}

	var provinces []Province
	for rows.Next() {
		var province Province
		err = rows.Scan(&province.ID, &province.Code, &province.Name, &province.NameEng, &province.GeoID)
		if err != nil {
			return nil, err
		}
		provinces = append(provinces, province)
	}
	rows.Close()

	return provinces, err
}

func GetDistrictsByProvince(province Province) ([]Amphur, error) {

	database, err := sql.Open("sqlite3", "data/th_address.db")
	checkErr(err)
	defer database.Close()

	//get amphur
	command := "SELECT * FROM amphures WHERE province_id=?"
	rows, err := database.Query(command, province.ID)
	if err != nil {
		return nil, err
	}

	var amphures []Amphur
	for rows.Next(){
		var amphur Amphur
		err = rows.Scan(&amphur.ID, &amphur.Code, &amphur.Name, &amphur.NameEng, &amphur.GeoID, &amphur.ProvinceID)
		if err != nil {
			return nil, err
		}
		amphures = append(amphures, amphur)
	}

	return amphures, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
