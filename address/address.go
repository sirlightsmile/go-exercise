package address

//	- NewAddress(subDistrict, district, province, zipcode) - return new Address instance
//	- GetProvinces() - return list of all provinces
//  - GetDistrictsByProvince(province) - return list of all districts in a specified province
//  - GetZipcodesByDistrict(district) - return list of all zipcodes in a specified district
//
//  - create Address method:
//	- Validate() - return true if the address is valid

func NewAddress(db *QueryInterface, subDistrictName string, districtName string, provinceName string, zipcode string) Address {

	subdistict, _ := GetSubDistrictByName(db, subDistrictName)
	district, _ := GetAmphurByName(db, districtName)
	province, _ := GetProvinceByName(db, provinceName)
	zipCode, _ := GetZipCodeModelByZipCode(db, zipcode)

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

func GetProvinces(db *QueryInterface) ([]Province, error) {
	rows, err := db.Query("SELECT * FROM provinces")
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

func GetDistrictsByProvince(db *QueryInterface, provinceName string) ([]Amphur, error) {

	query := `
		SELECT x.* FROM amphures x
		INNER JOIN provinces y ON x.province_id = y.province_id
		WHERE UPPER(y.province_name_eng) = UPPER(?)
	`

	rows, err := db.Query(query, provinceName)
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

func GetZipcodesByDistrict(db *QueryInterface, districtName string) ([]ZipCode, error) {

	query := `
		SELECT * FROM zipcodes z
		WHERE district_code COLLATE NOCASE IN 
		(SELECT district_code FROM districts WHERE amphur_id IN (SELECT amphur_id FROM amphures WHERE UPPER(amphur_name_eng) = UPPER(?)))
		GROUP BY z.zipcode COLLATE NOCASE
	`
	rows, err := db.Query(query, districtName)
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

func GetProvinceByName(db *QueryInterface, name string) (Province, error) {

	command := "SELECT * FROM provinces WHERE province_name = ? OR UPPER(province_name_eng) = UPPER(?)"
	row := db.QueryRow(command, name, name)
	var province Province
	err := row.Scan(&province.ID, &province.Code, &province.Name, &province.NameEng, &province.GeoID)
	if err != nil {
		return Province{}, err
	}
	return province, nil
}

func GetSubDistrictByName(db *QueryInterface, name string) (SubDistrict, error) {
	command := "SELECT * FROM districts WHERE district_name = ? OR UPPER(district_name_eng) = UPPER(?)"
	row := db.QueryRow(command, name)
	var subDistrict SubDistrict
	err := row.Scan(&subDistrict.ID, &subDistrict.Code, &subDistrict.Name, &subDistrict.NameEng, &subDistrict.AmphurID, &subDistrict.ProvinceID, &subDistrict.GeoID)
	if err != nil {
		return SubDistrict{}, err
	}
	return subDistrict, nil
}

func GetAmphurByName(db *QueryInterface, name string) (Amphur, error) {

	command := "SELECT * FROM amphures WHERE amphur_name = ? OR UPPER(amphur_name_eng) = UPPER(?)"
	row := db.QueryRow(command, name, name)
	var amphur Amphur
	err := row.Scan(&amphur.ID, &amphur.Code, &amphur.Name, &amphur.NameEng, &amphur.GeoID, &amphur.ProvinceID)
	if err != nil {
		return Amphur{}, err
	}
	return amphur, nil
}

func GetZipCodeModelByZipCode(db *QueryInterface, zipcode string) (ZipCode, error) {

	command := "SELECT * FROM zipcodes WHERE zipcode = ? COLLATE NOCASE"
	row := db.QueryRow(command, zipcode)
	var result ZipCode
	err := row.Scan(&result.ID, &result.SubDistrict, &result.ZipCode)
	if err != nil {
		return ZipCode{}, err
	}
	return result, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
