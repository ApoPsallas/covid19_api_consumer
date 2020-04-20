package main

import (
	"log"
	"net/http"

	app_http "github.com/ApoPsallas/simpleGoMicroservice/internal/http"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/affected_countries", app_http.AffectedCountriesHandler)
	err := http.ListenAndServe(":5000", r)
	if err != nil {
		log.Panic(err)
	}

}
