package address

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
