package api

import (
	"fmt"
	"net/http"
	"smile/api/addressAPI"
	"smile/repository"
)

func Init(db repository.QueryInterface, port string) {

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
			apiHandler.GetHandler(db, w, r)
		})
	}

	http.ListenAndServe(port, nil)
}
