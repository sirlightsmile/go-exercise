package api

import (
	"net/http"
	"smile/address"
)

type Handler interface {
	GetHandler(qi *address.AddressManager, w http.ResponseWriter, r *http.Request)
	GetAPIName() string
}
