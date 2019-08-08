package api

import (
	"encoding/json"
	"net/http"
	"smile/address"
	"smile/repository"
)

type NewAddress struct{
	handler Handler
}

func (api *NewAddress) GetHandler(qi repository.QueryInterface, w http.ResponseWriter, r *http.Request) {
	var task struct {
		Province    string
		District    string
		SubDistrict string
		ZipCode     string
	}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	newAddress := address.NewAddress(qi, task.SubDistrict, task.District, task.Province, task.ZipCode)
	json.NewEncoder(w).Encode(newAddress)
}

func (api *NewAddress) GetAPIName() string {
	return "/newAddress"
}
