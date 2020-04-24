package http

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/ApoPsallas/simpleGoMicroservice/internal/app"
)

type Handlers struct {
	Service app.API_service
}

//AffectedCountriesHandler handles incoming HTTP request
func (handlers Handlers) AffectedCountriesHandler(w http.ResponseWriter, r *http.Request) {

	affectedCountries := app.AffectedCountries{}
	response, err := handlers.Service.GetAffectedCountriesClient()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.Unmarshal(response, &affectedCountries)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Wrong structure of JSON response")
		return
	}
	sort.Strings(affectedCountries.AffectedCountries)
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(affectedCountries)

}
