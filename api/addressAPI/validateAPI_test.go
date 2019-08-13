package addressAPI

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidaionAPI(t *testing.T) {

	t.Run("Validate api test", func(t *testing.T) {

		testApi := &Validate{}

		req, err := http.NewRequest("GET", testApi.GetAPIName(), strings.NewReader(`{"SubDistrict":{"ID":1,"Code":"100101","Name":"พระบรมมหาราชวัง","NameEng":"Phra Borom Maha Ratchawang","GeoID":2,"AmphurID":1,"ProvinceID":1},"District":{"ID":1,"Code":"1001","Name":"เขตพระนคร   ","NameEng":"Khet Phra Nakhon","GeoID":2,"ProvinceID":1},"Province":{"ID":1,"Code":"10","Name":"กรุงเทพมหานคร   ","NameEng":"Bangkok","GeoID":2},"ZipCode":{"ID":1,"SubDistrict":"100101","ZipCode":"10200"}}`))
		if err != nil {
			t.Fatal(err)
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			testApi.GetHandler(testAddressManager, w, r)
		})

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `true` + "\n" //encode auto new line. response need \n
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
