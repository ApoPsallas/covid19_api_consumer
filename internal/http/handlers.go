package http

import (
	"encoding/json"
	"net/http"

	"github.com/ApoPsallas/simpleGoMicroservice/internal/app"
)

//AffectedCountriesHandler handles incoming HTTP request
func AffectedCountriesHandler(w http.ResponseWriter, r *http.Request) {

	affected, err := app.GetAffectedCountries()
	if err != nil {
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(affected)

}
