package api

import (
	"fmt"
	"net/http"
	"smile/repository"
)

func Init(db repository.QueryInterface, port string) {

	apiList := []Handler{
		&GetProvinceAPI{},
		&GetDistrictsByProvince{},
		&GetZipcodesByDistrict{},
		&NewAddress{},
		&Validate{},
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
