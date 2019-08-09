package addressAPI

import (
	"encoding/json"
	"net/http"
	"smile/address"
)

type GetProvinceAPI struct{}

func (api *GetProvinceAPI) GetHandler(ai address.AddressModel, w http.ResponseWriter, r *http.Request) {

	provinces, _ := ai.GetProvinces()
	json.NewEncoder(w).Encode(provinces)
}

func (api *GetProvinceAPI) GetAPIName() string {
	return "/getProvinces"
}
