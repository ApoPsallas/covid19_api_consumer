package main

import (
	"log"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/ApoPsallas/covid19_api_consumer/internal/app"
	app_http "github.com/ApoPsallas/covid19_api_consumer/internal/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)

func createRouter() *mux.Router {
	router := mux.NewRouter()
	service := app.RapidapiService{Client: http.DefaultClient}
	handlers := app_http.Handlers{Service: service}
	router.HandleFunc("/affected_countries", handlers.AffectedCountriesHandler)
	return router

}

func main() {

	err := godotenv.Load(Root + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := createRouter()
	err = http.ListenAndServe(":5000", router)
	if err != nil {
		log.Panic(err)
	}

}
