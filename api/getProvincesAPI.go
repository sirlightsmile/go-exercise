package api

import (
	"encoding/json"
	"net/http"
	"smile/address"
	"smile/repository"
)

type GetProvinceAPI struct{}

func (api *GetProvinceAPI) GetHandler(qi repository.QueryInterface, w http.ResponseWriter, r *http.Request) {
	provinces, err := address.GetProvinces(qi)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(provinces)
}

func (api *GetProvinceAPI) GetAPIName() string {
	return "/getProvince"
}
