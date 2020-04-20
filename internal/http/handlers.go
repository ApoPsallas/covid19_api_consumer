package http

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/ApoPsallas/simpleGoMicroservice/internal/app"
)

//AffectedCountriesHandler handles incoming HTTP request
func AffectedCountriesHandler(w http.ResponseWriter, r *http.Request) {

	//TODO inject api_service
	api := app.API{Client: http.DefaultClient}
	ac := app.AffectedCountries{}
	affected, err := api.GetAffectedCountriesClient()
	if err != nil {
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	_ = json.Unmarshal(affected, &ac)
	sort.Strings(ac.AffectedCountries)
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(ac)

}