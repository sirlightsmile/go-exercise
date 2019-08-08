package api

import (
	"net/http"
	"smile/repository"
)

func Init(db repository.QueryInterface, port string) {

	apiList := []Handler{}
	apiList = append(apiList, GetDistrictsByProvince{})

	for _, v := range apiList {
		http.HandleFunc(v.GetAPIName(), func(w http.ResponseWriter, r *http.Request) {
			v.GetHandler(db, w, r)
		})
	}

	/*
		provinceAPI := GetProvinceAPI{}
		http.HandleFunc(provinceAPI.GetAPIName(), func(w http.ResponseWriter, r *http.Request) {
			provinceAPI.GetHandler(db, w, r)
		})
	*/
	http.ListenAndServe(":4000", nil)
}
