package address

import "smile/repository"

//	- NewAddress(subDistrict, district, province, zipcode) - return new Address instance
//	- GetProvinces() - return list of all provinces
//  - GetDistrictsByProvince(province) - return list of all districts in a specified province
//  - GetZipcodesByDistrict(district) - return list of all zipcodes in a specified district
//
//  - create Address method:
//	- Validate() - return true if the address is valid

func NewAddress(qi repository.QueryInterface, subDistrictName string, districtName string, provinceName string, zipcode string) Address {

	subdistict, _ := GetSubDistrictByName(qi, subDistrictName)
	district, _ := GetAmphurByName(qi, districtName)
	province, _ := GetProvinceByName(qi, provinceName)
	zipCode, _ := GetZipCodeModelByZipCode(qi, zipcode)

	address := Address{
		SubDistrict: subdistict,
		District:    district,
		Province:    province,
		ZipCode:     zipCode,
	}

	return address
}

func Validate(address Address) bool {
	return address.ZipCode.SubDistrict == address.SubDistrict.Code &&
		address.SubDistrict.AmphurID == address.District.ID &&
		address.District.ProvinceID == address.Province.ID
}

func GetProvinces(qi repository.QueryInterface) ([]Province, error) {
	rows, err := qi.Query("SELECT * FROM provinces")
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

func GetDistrictsByProvince(qi repository.QueryInterface, provinceName string) ([]Amphur, error) {

	query := `
		SELECT x.* FROM amphures x
		INNER JOIN provinces y ON x.province_id = y.province_id
		WHERE UPPER(y.province_name_eng) = UPPER(?)
	`

	rows, err := qi.Query(query, provinceName)
	if err != nil {
		return nil, err
	}

	var amphures []Amphur
	for rows.Next() {
		var amphur Amphur
		err = rows.Scan(&amphur.ID, &amphur.Code, &amphur.Name, &amphur.NameEng, &amphur.GeoID, &amphur.ProvinceID)
		if err != nil {
			return nil, err
		}
		amphures = append(amphures, amphur)
	}

	return amphures, nil
}

func GetZipcodesByDistrict(qi repository.QueryInterface, districtName string) ([]ZipCode, error) {

	query := `
		SELECT * FROM zipcodes z
		WHERE district_code COLLATE NOCASE IN 
		(SELECT district_code FROM districts WHERE amphur_id IN (SELECT amphur_id FROM amphures WHERE UPPER(amphur_name_eng) = UPPER(?)))
	`
	rows, err := qi.Query(query, districtName)
	if err != nil {
		return nil, err
	}

	var zipcodes []ZipCode
	for rows.Next() {
		var zipcode ZipCode
		err := rows.Scan(&zipcode.ID, &zipcode.SubDistrict, &zipcode.ZipCode)
		if err != nil {
			return nil, err
		}
		zipcodes = append(zipcodes, zipcode)
	}

	return zipcodes, nil
}

func GetProvinceByName(qi repository.QueryInterface, name string) (Province, error) {

	command := "SELECT * FROM provinces WHERE TRIM(province_name) = TRIM(?) COLLATE NOCASE OR UPPER(province_name_eng) = UPPER(?)"
	row := qi.QueryRow(command, name, name)
	var province Province
	err := row.Scan(&province.ID, &province.Code, &province.Name, &province.NameEng, &province.GeoID)
	if err != nil {
		return Province{}, err
	}
	return province, nil
}

func GetSubDistrictByName(qi repository.QueryInterface, name string) (SubDistrict, error) {
	command := "SELECT * FROM districts WHERE TRIM(district_name) = TRIM(?) COLLATE NOCASE OR UPPER(district_name_eng) = UPPER(?)"
	row := qi.QueryRow(command, name, name)
	var subDistrict SubDistrict
	err := row.Scan(&subDistrict.ID, &subDistrict.Code, &subDistrict.Name, &subDistrict.NameEng, &subDistrict.AmphurID, &subDistrict.ProvinceID, &subDistrict.GeoID)
	if err != nil {
		return SubDistrict{}, err
	}
	return subDistrict, nil
}

func GetAmphurByName(qi repository.QueryInterface, name string) (Amphur, error) {

	command := "SELECT * FROM amphures WHERE TRIM(amphur_name) = TRIM(?) COLLATE NOCASE OR UPPER(amphur_name_eng) = UPPER(?)"
	row := qi.QueryRow(command, name, name)
	var amphur Amphur
	err := row.Scan(&amphur.ID, &amphur.Code, &amphur.Name, &amphur.NameEng, &amphur.GeoID, &amphur.ProvinceID)
	if err != nil {
		return Amphur{}, err
	}
	return amphur, nil
}

func GetZipCodeModelByZipCode(qi repository.QueryInterface, zipcode string) (ZipCode, error) {

	command := "SELECT * FROM zipcodes WHERE zipcode = ? COLLATE NOCASE"
	row := qi.QueryRow(command, zipcode)
	var result ZipCode
	err := row.Scan(&result.ID, &result.SubDistrict, &result.ZipCode)
	if err != nil {
		return ZipCode{}, err
	}
	return result, nil
}
