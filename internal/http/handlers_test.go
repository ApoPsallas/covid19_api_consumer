package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAffectedCountriesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/affected_countries", nil)
	assert.Nil(t, err)
	assert.NotNil(t, req)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(AffectedCountriesHandler)
	handler.ServeHTTP(r, req)
	assert.Equal(t, r.Code, http.StatusOK)

}
