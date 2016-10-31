package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"io"
)

func Handlers() *mux.Router{

	r:=mux.NewRouter()
	r.HandleFunc("/healthcheck", HealthCheckHandler).Methods("GET")
	return r
}


func HealthCheckHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	io.WriteString(w,`{"alive":"ok"}`)

}