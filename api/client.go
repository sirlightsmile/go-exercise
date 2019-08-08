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
		fmt.Println("API name : ", v.GetAPIName())
		http.HandleFunc(v.GetAPIName(), func(w http.ResponseWriter, r *http.Request) {
			v.GetHandler(db, w, r)
		})
	}

	http.ListenAndServe(":4000", nil)
}
