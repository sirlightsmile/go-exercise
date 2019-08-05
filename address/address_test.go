package address

import (
	"os"
	"path/filepath"
	"reflect"
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
var absPath string

func TestMain(m *testing.M) {
	absPath, _ = filepath.Abs("../data/th_address.db")
	code := m.Run()
	os.Exit(code)
}

func TestNewAddress(t *testing.T) {
	db, _ := ConnectSqlDB(absPath)

	t.Run("New address test", func(t *testing.T) {
		expected := Address{
			Province: Province{
				ID:      1,
				Code:    "10",
				Name:    "กรุงเทพมหานคร   ",
				NameEng: "Bangkok",
				GeoID:   2,
			},
			District: Amphur{
				ID:         1,
				Code:       "1001",
				Name:       "เขตพระนคร   ",
				NameEng:    "Khet Phra Nakhon",
				GeoID:      2,
				ProvinceID: 1,
			},
			SubDistrict: SubDistrict{
				ID:         1,
				Code:       "100101",
				Name:       "พระบรมมหาราชวัง",
				NameEng:    "Phra Borom Maha Ratchawang",
				GeoID:      2,
				AmphurID:   1,
				ProvinceID: 1,
			},
			ZipCode: ZipCode{
				ID:          1,
				SubDistrict: "100101",
				ZipCode:     "10200",
			},
		}

		falseExpected := Address{
			Province:    Province{},
			District:    Amphur{},
			SubDistrict: SubDistrict{},
			ZipCode:     ZipCode{},
		}

		result := NewAddress(db, "Phra Borom Maha Ratchawang", "Khet Phra Nakhon", "Bangkok", "10200")
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Failed, expected : \n\n%#v\n\nreality : \n\n%#v", expected, result)
		}

		result = NewAddress(db, "?", "?", "?", "?")
		if !reflect.DeepEqual(falseExpected, result) {
			t.Errorf("Failed, expected : \n\n%#v\n\nreality : \n\n%#v", expected, result)
		}
	})
}

func TestGetProvinces(t *testing.T) {

	db, _ := ConnectSqlDB(absPath)

	t.Run("Get provinces test", func(t *testing.T) {
		expected := 77
		result, err := GetProvinces(db)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		reality := len(result)
		if expected != reality {
			t.Errorf("Failed, expected : %d reality : %d", expected, reality)
		}
	})
}

func TestGetDistrictsByProvince(t *testing.T) {

	db, _ := ConnectSqlDB(absPath)

	t.Run("Get district by province test", func(t *testing.T) {

		expectedResult := []string{
			"1101",
			"1102",
			"1103",
			"1104",
			"1105",
			"1106",
		}

		result, err := GetDistrictsByProvince(db, "Samut Prakan")
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		for i, element := range result {
			expected := expectedResult[i]
			if element.Code != expected {
				t.Errorf("Failed, expected : %s reality : %s", expected, element.Code)
			}
		}
	})
}

func TestGetZipcodesByDistrict(t *testing.T) {

	db, _ := ConnectSqlDB(absPath)

	t.Run("Get zipcode by district test", func(t *testing.T) {

		expectedResult := []string{
			"100401",
			"100402",
			"100403",
			"100404",
			"100405",
		}

		result, err := GetZipcodesByDistrict(db, "Khet Bang Rak")
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		for i, element := range result {
			expected := expectedResult[i]
			if element.SubDistrict != expected {
				t.Errorf("Failed, expected : %s reality : %s", expected, element.SubDistrict)
			}
		}
	})
}
