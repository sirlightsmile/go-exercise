package addressAPI

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"smile/address"
	"smile/repository"
	"strings"
	"testing"
)

func TestNewAddressAPI(t *testing.T) {
	absPath, _ := filepath.Abs("../../data/th_address.db")
	db, err := repository.ConnectSqlDB(absPath)
	if err != nil {
		panic(err)
	}
	am := address.Init(db)

	t.Run("Get new address api test", func(t *testing.T) {

		testApi := &NewAddress{}

		req, err := http.NewRequest("GET", testApi.GetAPIName(), strings.NewReader(`{"Province":"Bangkok","District":"Khet Phra Nakhon","SubDistrict":"Phra Borom Maha Ratchawang","Zipcode" : "10200"}`))

		if err != nil {
			t.Fatal(err)
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			testApi.GetHandler(am, w, r)
		})

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `{"SubDistrict":{"ID":1,"Code":"100101","Name":"พระบรมมหาราชวัง","NameEng":"Phra Borom Maha Ratchawang","GeoID":2,"AmphurID":1,"ProvinceID":1},"District":{"ID":1,"Code":"1001","Name":"เขตพระนคร   ","NameEng":"Khet Phra Nakhon","GeoID":2,"ProvinceID":1},"Province":{"ID":1,"Code":"10","Name":"กรุงเทพมหานคร   ","NameEng":"Bangkok","GeoID":2},"ZipCode":{"ID":1,"SubDistrict":"100101","ZipCode":"10200"}}` + "\n" //encode auto new line. response need \n
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}

		//bad request test
		badreq, err := http.NewRequest("GET", testApi.GetAPIName(), strings.NewReader(`bad`))
		if err != nil {
			t.Fatal(err)
		}

		rr = httptest.NewRecorder()
		handler.ServeHTTP(rr, badreq)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		// Check the response body is what we expect.
		expected = `bad request` + "\n" //encode auto new line. response need \n
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
