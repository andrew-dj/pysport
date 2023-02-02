package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run(addr string, pathToPWA string) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/races", getRaces).Methods("GET")
	if pathToPWA != "" {
		r.PathPrefix("/").Handler(http.FileServer(http.Dir(pathToPWA)))
	}
	log.Fatal(http.ListenAndServe(addr, r))
}
