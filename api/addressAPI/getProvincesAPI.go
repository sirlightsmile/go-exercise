package addressAPI

import (
	"encoding/json"
	"net/http"
	"smile/address"
)

type GetProvinceAPI struct{}

func (api *GetProvinceAPI) GetHandler(am *address.AddressManager, w http.ResponseWriter, r *http.Request) {

	provinces, _ := am.GetProvinces()
	json.NewEncoder(w).Encode(provinces)
}

func (api *GetProvinceAPI) GetAPIName() string {
	return "/getProvinces"
}
