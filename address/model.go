package address

// 	subDistrict (districts table)
//		- district (amphures table)
//		- province (provinces table)
//		- zipcode (zipcodes table)

type Address struct {
	SubDistrict SubDistrict
	District    Amphur
	Province    Province
	ZipCodes    ZipCode
}

type Amphur struct {
	ID         int
	Code       string
	Name       string
	NameEng    string
	ProvinceID int
	GeoID      int
}

type SubDistrict struct {
	ID         int
	Code       string
	Name       string
	NameEng    string
	AmphurID   int
	ProvinceID int
	GeoID      int
}

type Province struct {
	ID      int
	Code    int
	Name    string
	NameEng string
	GeoID   int
}

type ZipCode struct {
	ID              int
	SubDistrictCode int
	ZipCode         string
}
