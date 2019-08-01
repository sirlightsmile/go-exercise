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
	"strings"
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func Init() error{
	var err error
	database, err = sql.Open("sqlite3", "data/th_address.db")
	checkErr(err)
	return database.Ping();
}

func NewAddress(subDistrict string, district string, province string, zipcode string) (Address, error){
	var address Address
	var err error
	address.Province, err = GetProvinceByName(province)
	if(err != nil){
		return Address{}, err
	}
	address.SubDistrict, err = GetSubDistrictByName(subDistrict)
	if(err != nil){
		return Address{}, err
	}
	address.District, err = GetAmphurByName(district)
	if(err != nil){
		return Address{}, err
	}
	address.ZipCodes, err = GetZipCodeModelByZipCode(zipcode)
	if(err != nil){
		return Address{}, err
	}

	return address, nil
}

func GetProvinces() ([]Province, error) {
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

	return provinces, nil
}

func GetDistrictsByProvince(provinceName string) ([]Amphur, error) {
	
	province, err := GetProvinceByName(provinceName)
	if err != nil {
		return nil, err
	}

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

	return amphures, nil
}

func GetZipcodesByDistrict(districtName string) (ZipCode, error) {

	amphur, err := GetAmphurByName(districtName)
	if err != nil {
		return ZipCode{}, err
	}

	command := "SELECT * FROM districts WHERE amphur_id=?"
	row := database.QueryRow(command, amphur.ID)
	var subDistrict SubDistrict
	err = row.Scan(&subDistrict.ID, &subDistrict.Code, &subDistrict.Name, &subDistrict.NameEng, &subDistrict.AmphurID, &subDistrict.ProvinceID, &subDistrict.GeoID)
	if err != nil {
		return ZipCode{}, err
	}

	command = "SELECT * FROM zipcodes WHERE district_code=?"
	row = database.QueryRow(command, subDistrict.Code)
	var zipcode ZipCode
	err = row.Scan(&zipcode.ID, &zipcode.SubDistrictCode, &zipcode.ZipCode)
	if err != nil {
		return ZipCode{}, err
	}

	return zipcode, nil
}

func GetProvinceByName(name string) (Province, error){

	name = strings.ToUpper(name)
	command := "SELECT * FROM provinces WHERE UPPER(province_name)=? OR UPPER(province_name_eng)=?"
	row := database.QueryRow(command, name, name)
	var province Province
	err := row.Scan(&province.ID, &province.Code, &province.Name, &province.NameEng, &province.GeoID)
	if(err != nil){
		return Province{}, err
	}
	return province, nil
}

func GetSubDistrictByName(name string) (SubDistrict, error){

	name = strings.ToUpper(name);
	command := "SELECT * FROM districts WHERE UPPER(district_name)=? OR UPPER(district_name_eng)=?"
	row := database.QueryRow(command, name, name)
	var subDistrict SubDistrict
	err := row.Scan(&subDistrict.ID, &subDistrict.Code, &subDistrict.Name, &subDistrict.NameEng, &subDistrict.AmphurID, &subDistrict.ProvinceID, &subDistrict.GeoID)
	if(err != nil){
		return SubDistrict{}, err
	}
	return subDistrict, nil
}

func GetAmphurByName(name string) (Amphur, error){
	
	name = strings.ToUpper(name)
	command := "SELECT * FROM amphures WHERE UPPER(amphur_name)=? OR UPPER(amphur_name_eng)=?"
	row := database.QueryRow(command, name, name)
	var amphur Amphur
	err := row.Scan(&amphur.ID, &amphur.Code, &amphur.Name, &amphur.NameEng, &amphur.GeoID, &amphur.ProvinceID)
	if(err != nil){
		return Amphur{}, err
	}
	return amphur, nil
}

func GetZipCodeModelByZipCode(zipcode string) (ZipCode, error){
	
	command := "SELECT * FROM zipcodes WHERE zipcode = ?"
	row := database.QueryRow(command, zipcode)
	var result ZipCode
	err := row.Scan(&result.ID, &result.SubDistrictCode, &result.ZipCode)
	if(err != nil){
		return ZipCode{}, err
	}
	return result, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
