package api

import (
	"encoding/json"
	"net/http"
	"smile/address"
	"smile/repository"
)

type GetZipcodesByDistrict struct{}

func (api *GetZipcodesByDistrict) GetHandler(qi repository.QueryInterface, w http.ResponseWriter, r *http.Request) {
	type Task struct {
		District string
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	zipcodes, err := address.GetZipcodesByDistrict(qi, task.District)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(zipcodes)
}

func (api *GetZipcodesByDistrict) GetAPIName() string {
	return "/getZipcodesByDistrict"
}
