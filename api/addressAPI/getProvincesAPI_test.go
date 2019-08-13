package addressAPI

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"smile/address"
	"testing"
)

func TestGetProvinceAPI(t *testing.T) {

	t.Run("Get province api test", func(t *testing.T) {
		testApi := &GetProvinceAPI{}

		req, err := http.NewRequest("GET", testApi.GetAPIName(), nil)
		if err != nil {
			t.Fatal(err)
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			testApi.GetHandler(TestManager, w, r)
		})

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		var provinces []address.Province
		err = json.Unmarshal([]byte(rr.Body.String()), &provinces)
		if err != nil {
			t.Fatal(err)
		}
		expected := 77
		if len(provinces) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				len(provinces), expected)
		}
	})
}
