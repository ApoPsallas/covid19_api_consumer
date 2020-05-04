package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/ApoPsallas/covid19_api_consumer/internal/app"
	"github.com/ApoPsallas/covid19_api_consumer/internal/app/repository"
	app_http "github.com/ApoPsallas/covid19_api_consumer/internal/http"
	"github.com/go-redis/redis/v7"
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
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	cache := repository.CacheMapper{Client: *client}
	service := app.RapidapiService{Client: http.DefaultClient, CacheMapper: cache}
	handlers := app_http.Handlers{Service: service}
	router.HandleFunc("/affected_countries", handlers.AffectedCountriesHandler)
	return router

}

func main() {

	fmt.Println("App started.")
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
