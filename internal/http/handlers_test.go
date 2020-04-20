package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
)

func TestAffectedCountriesHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/affected_countries", nil)
	resp := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/affected_countries", AffectedCountriesHandler)
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)

}
