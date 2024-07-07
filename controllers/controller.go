package controllers

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	configs "bittrack-backend/models/configs"
)

type App struct {
	Router *mux.Router
}

var config configs.Config

// InitConfig get variable from config file .env
func InitConfig(
	apiPortConfig string,
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
	}
}

func (a *App) Initialize() {
	// var err error
	// DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	// if _, err := os.Stat(config.FileStorage); os.IsNotExist(err) {
	// 	os.Mkdir(config.FileStorage, 0700)
	// }

	// a.DB, err = sql.Open("postgres", DBURI)
	// if err != nil {
	// 	log.Fatal("Cannot connect to database : ", err)
	// } else {
	// 	log.Println("We are connected to the database ", DbName)
	// }
	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}