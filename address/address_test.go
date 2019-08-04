package address

import (
	"testing"
)

//	- NewAddress(subDistrict, district, province, zipcode) - return new Address instance
//	- GetProvinces() - return list of all provinces
//  - GetDistrictsByProvince(province) - return list of all districts in a specified province
//  - GetZipcodesByDistrict(district) - return list of all zipcodes in a specified district
//
//  - create Address method:
//	- Validate() - return true if the address is valid
/*
var dbPath string

func TestMain(m *testing.M) {
	dbPath, _ = filepath.Abs("../data/th_address.db")
	code := m.Run()
	os.Exit(code)
}
*/
var dbPath string

func TestGetProvinces(t *testing.T) {
	db, _ := ConnectSqlDB(dbPath)

	t.Run("Get provinces test", func(t *testing.T) {
		expected := 77
		result, err := GetProvinces(db)
		if err != nil {
			t.Errorf("Error: ", err)
		}
		reality := len(result)
		if expected != reality {
			t.Errorf("Failed, expected %d to be %d", expected, reality)
		}
	})
}

func TestGetDistrictsByProvince(t *testing.T) {
	db, _ := ConnectSqlDB(dbPath)

	t.Run("Get district by province test", func(t *testing.T) {
		expected := 6
		result, err := GetDistrictsByProvince(db, "Samut Prakan")
		if err != nil {
			t.Errorf("Error: ", err)
		}
		reality := len(result)
		if expected != reality {
			t.Errorf("Failed, expected : %d actual : %d", expected, reality)
		}
	})
}
