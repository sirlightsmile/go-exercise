package address

import "smile/repository"

type AddressInterface interface {
	NewAddress(subDistrictName string, districtName string, provinceName string, zipcode string) Address
	GetProvinces() ([]Province, error)
	GetDistrictsByProvince(provinceName string) ([]Amphur, error)
	GetZipcodesByDistrict(districtName string) ([]ZipCode, error)
	Validate(address Address) bool
}

type AddressModel struct {
	repo repository.QueryInterface
}

type Address struct {
	SubDistrict SubDistrict
	District    Amphur
	Province    Province
	ZipCode     ZipCode
}

type Amphur struct {
	ID         int
	Code       string
	Name       string
	NameEng    string
	GeoID      int
	ProvinceID int
}

type SubDistrict struct {
	ID         int
	Code       string
	Name       string
	NameEng    string
	GeoID      int
	AmphurID   int
	ProvinceID int
}

type Province struct {
	ID      int
	Code    string
	Name    string
	NameEng string
	GeoID   int
}

type ZipCode struct {
	ID          int
	SubDistrict string
	ZipCode     string
}
