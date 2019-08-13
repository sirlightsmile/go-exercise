package api

import (
	"fmt"
	"net/http"
	"smile/address"
	"smile/api/addressAPI"
)

func Init(am *address.AddressManager, port string) {

	apiList := []Handler{
		&addressAPI.GetProvinceAPI{},
		&addressAPI.GetDistrictsByProvince{},
		&addressAPI.GetZipcodesByDistrict{},
		&addressAPI.NewAddress{},
		&addressAPI.Validate{},
	}

	for _, v := range apiList {
		apiHandler := v
		fmt.Println("API name : ", v.GetAPIName())
		http.HandleFunc(apiHandler.GetAPIName(), func(w http.ResponseWriter, r *http.Request) {
			apiHandler.GetHandler(am, w, r)
		})
	}

	http.ListenAndServe(port, nil)
}
