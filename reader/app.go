package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Reading environment failed.")
	}

	if !fileExists(os.Getenv("APP_LOG_FILE")) {
		createFile(os.Getenv("APP_LOG_FILE"), "")
	}
	f, err := os.OpenFile(os.Getenv("APP_LOG_FILE"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Opening log file failed: %s", err)
	}

	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Printf("starting reader")
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.getRoot).Methods("GET")
	a.Router.HandleFunc("/health", a.getHealth).Methods("GET")
}

func (a *App) Run() {

	headers := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{fmt.Sprintf("http://%s", os.Getenv("ALLOWED_ORIGINS"))})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodOptions, http.MethodConnect, http.MethodPost})
	maxAge := handlers.MaxAge(60)

	address := fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT"))
	server := &http.Server{
		Addr:    address,
		Handler: handlers.CORS(headers, origins, methods, maxAge)(a.Router),
	}

	log.Printf("starting log-output server in %s.", address)
	log.Printf("Version: %s , build: %s", Version, Build)
	log.Printf("Allowed origins: %s", os.Getenv("ALLOWED_ORIGINS"))
	log.Fatal(server.ListenAndServe())
}
