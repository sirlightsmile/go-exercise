package address

type Address struct {
	SubDistrict string
	District    string
	Province    string
	Zipcodes    string
}

type Province struct {
	ProvinceID      uint
	ProvinceCode    uint
	ProvinceName    string
	ProvinceNameEng string
	GeoID           uint
}
