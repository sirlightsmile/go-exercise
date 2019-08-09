package addressAPI

import (
	"encoding/json"
	"net/http"
	"smile/address"
)

type NewAddress struct{}

func (api *NewAddress) GetHandler(ai address.AddressModel, w http.ResponseWriter, r *http.Request) {
	var task struct {
		Province    string `json:"Province"`
		District    string `json:"District"`
		SubDistrict string `json:"SubDistrict"`
		ZipCode     string `json:"Zipcode"`
	}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	newAddress := ai.NewAddress(task.SubDistrict, task.District, task.Province, task.ZipCode)
	json.NewEncoder(w).Encode(newAddress)
}

func (api *NewAddress) GetAPIName() string {
	return "/newAddress"
}
