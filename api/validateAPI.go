package api

import (
	"encoding/json"
	"net/http"
	"smile/address"
	"smile/repository"
)

type Validate struct {
}

func (api *Validate) GetHandler(qi repository.QueryInterface, w http.ResponseWriter, r *http.Request) {
	var task address.Address
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	valid := address.Validate(task)

	json.NewEncoder(w).Encode(valid)
}

func (api *Validate) GetAPIName() string {
	return "/validate"
}
