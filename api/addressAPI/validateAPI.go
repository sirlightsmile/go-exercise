package addressAPI

import (
	"encoding/json"
	"net/http"
	"smile/address"
)

type Validate struct{}

func (api *Validate) GetHandler(ai address.AddressModel, w http.ResponseWriter, r *http.Request) {
	var task address.Address
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	valid := ai.Validate(task)

	json.NewEncoder(w).Encode(valid)
}

func (api *Validate) GetAPIName() string {
	return "/validate"
}
