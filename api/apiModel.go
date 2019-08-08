package api

import (
	"net/http"
	"smile/repository"
)

type Handler interface {
	GetHandler(qi repository.QueryInterface, w http.ResponseWriter, r *http.Request)
	GetAPIName() string
}
