package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	_ "github.com/lib/pq"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

)

type App struct {
	Router *mux.Router
}

// InitConfig get variable from config file .env
func InitConfig(
	apiPortConfig,
) {
	config.APIPort = apiPortConfig

}

func (a *App) initializeRoutes() {
	// home page
	a.Get("/", home)
}

func (a *App) RunServer() {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Accept", "Authorization", "Accesstoken"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	log.Printf("\nServer starting on port %s", config.APIPort)
	log.Fatal(http.ListenAndServe((fmt.Sprintf(":%s", config.APIPort)), handlers.CORS(originsOk, headersOk, methodsOk)(a.Router)))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, this API works. Check documentation for usage\n"))
}

// reference --> https://github.com/mingrammer/go-todo-rest-api-example/blob/master/app/app.go

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Delete wraps the router for DELETE method
func (a *App) Options(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("OPTIONS")
}

// RequestService represents generalized form of services
type RequestService func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")

		if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
        }

		handler(a.DB, config, logs, w, r)
	}
}