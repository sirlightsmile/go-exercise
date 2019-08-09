package addressAPI

import (
	"encoding/json"
	"net/http"
	"smile/address"
)

type GetDistrictsByProvince struct{}

func (api *GetDistrictsByProvince) GetHandler(ai address.AddressModel, w http.ResponseWriter, r *http.Request) {
	type Task struct {
		Province string
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	districts, _ := ai.GetDistrictsByProvince(task.Province)

	json.NewEncoder(w).Encode(districts)
}

func (api *GetDistrictsByProvince) GetAPIName() string {
	return "/getDistrictByProvince"
}
