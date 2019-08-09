package addressAPI

import (
	"encoding/json"
	"net/http"
	"smile/address"
	"smile/repository"
)

type GetDistrictsByProvince struct{}

func (api *GetDistrictsByProvince) GetHandler(qi repository.QueryInterface, w http.ResponseWriter, r *http.Request) {
	type Task struct {
		Province string
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	districts, err := address.GetDistrictsByProvince(qi, task.Province)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(districts)
}

func (api *GetDistrictsByProvince) GetAPIName() string {
	return "/getDistrictByProvince"
}
