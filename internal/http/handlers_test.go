package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ApoPsallas/covid19_api_consumer/internal/app"
	"github.com/stretchr/testify/assert"
)

type MockApi struct {
	Response *app.AffectedCountries
	Err      error
}

func (api MockApi) GetAffectedCountries() (*app.AffectedCountries, error) {
	return api.Response, api.Err
}

func TestAffectedCountriesHandler(t *testing.T) {

	request := httptest.NewRequest("GET", "/affected_countries", nil)
	responseWriter := httptest.NewRecorder()

	affectedCountries := app.NewAffectedCountries([]string{"es", "gr", "it", "uk"}, "now")
	response, _ := json.Marshal(affectedCountries)

	api := MockApi{Response: affectedCountries, Err: nil}

	handlerStruct := Handlers{Service: api}

	handler := http.HandlerFunc(handlerStruct.AffectedCountriesHandler)

	handler.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, string(response), strings.TrimRight(responseWriter.Body.String(), "\n"))
	//there is a "\n" added from Encoder
}

func TestAffectedCountriesHandlerServiceUnavailable(t *testing.T) {

	request := httptest.NewRequest("GET", "/affected_countries", nil)
	responseWriter := httptest.NewRecorder()

	affectedCountries := app.AffectedCountries{}
	expected := "\"Service Unavailable\""

	api := MockApi{Response: &affectedCountries, Err: errors.New("Service Unavailable")}

	handlerStruct := Handlers{Service: api}

	handler := http.HandlerFunc(handlerStruct.AffectedCountriesHandler)

	handler.ServeHTTP(responseWriter, request)
	assert.Equal(t, http.StatusServiceUnavailable, responseWriter.Code)
	assert.Equal(t, expected, strings.TrimRight(responseWriter.Body.String(), "\n"))
	//there is a "\n" added from Encoder
}
