package http

import (
	"encoding/json"
	"net/http"

	"github.com/ApoPsallas/covid19_api_consumer/internal/app"
)

type Handlers struct {
	Service app.APIService
}

//AffectedCountriesHandler handles incoming HTTP request
func (handlers Handlers) AffectedCountriesHandler(w http.ResponseWriter, r *http.Request) {

	affectedCountries, err := handlers.Service.GetAffectedCountries()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(affectedCountries)

}
