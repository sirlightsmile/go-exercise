package address

type Address struct {
	SubDistrict District
	District    Amphur
	Province    Province
	ZipCodes    ZipCode
}

type Amphur struct {
	AmphurID      int
	AmphurCode    string
	AmphurName    string
	AmphurNameEng string
	ProvinceID    int
	GeoID         int
}

type District struct {
	DistrictID      int
	DistrictCode    string
	DistrictName    string
	DistrictNameEng string
	AmphurID        int
	ProvinceID      int
	GeoID           int
}

type Province struct {
	ProvinceID      int
	ProvinceCode    int
	ProvinceName    string
	ProvinceNameEng string
	GeoID           int
}

type ZipCode struct {
	ID           int
	DistrictCode int
	ZipCode      string
}
