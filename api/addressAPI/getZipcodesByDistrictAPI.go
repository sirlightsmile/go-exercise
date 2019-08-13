package addressAPI

import (
	"encoding/json"
	"net/http"
	"smile/address"
)

type GetZipcodesByDistrict struct{}

func (api *GetZipcodesByDistrict) GetHandler(am *address.AddressManager, w http.ResponseWriter, r *http.Request) {
	type Task struct {
		District string
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	zipcodes, _ := am.GetZipcodesByDistrict(task.District)

	json.NewEncoder(w).Encode(zipcodes)
}

func (api *GetZipcodesByDistrict) GetAPIName() string {
	return "/getZipcodesByDistrict"
}
