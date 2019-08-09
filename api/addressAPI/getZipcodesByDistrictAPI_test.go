package addressAPI

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"smile/address"
	"smile/repository"
	"strings"
	"testing"
)

func TestGetZipcodesByDistrictAPI(t *testing.T) {
	absPath, _ := filepath.Abs("../../data/th_address.db")
	db, err := repository.ConnectSqlDB(absPath)
	if err != nil {
		panic(err)
	}

	t.Run("test get zipcodes by district api test", func(t *testing.T) {
		testApi := &GetZipcodesByDistrict{}

		req, err := http.NewRequest("GET", testApi.GetAPIName(), strings.NewReader(`{"District" : "Khet Dusit"}`))
		if err != nil {
			t.Fatal(err)
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			testApi.GetHandler(db, w, r)
		})

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		var zipcodes []address.ZipCode
		err = json.Unmarshal([]byte(rr.Body.String()), &zipcodes)
		if err != nil {
			t.Fatal(err)
		}
		expectedCount := 5
		if len(zipcodes) != expectedCount {
			t.Errorf("handler returned unexpected body: got %v want %v",
				len(zipcodes), expectedCount)
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
		expected := `bad request` + "\n" //encode auto new line. response need \n
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
