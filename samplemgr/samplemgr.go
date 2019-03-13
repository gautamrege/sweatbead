package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/negroni"
)

var (
	db *sqlx.DB
)

func main() {

	// Initialize db
	err := initDB()
	if err != nil {
		panic(err)
	}

	// mux router
	router := initRouter()

	// init web server
	server := negroni.Classic()
	server.UseHandler(router)

	port := 33001 // This should be changed to the service port number via argument or environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)

}

/* Initialize the database */
func initDB() (err error) {
	db, err := sqlx.Connect("sqlite3", "__deleteme.db")
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Duration(30) * time.Minute)

	return
}

/* The routing mechanism. Mux helps us define handler functions and the access methods */
func initRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	return
}

type Response struct {
	Message string `json:"message"`
}

/* Each handler will then respond with marshalled Json data */
func pingHandler(rw http.ResponseWriter, req *http.Request) {
	response := Response{Message: "pong"}

	status := http.StatusOK
	respBytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
