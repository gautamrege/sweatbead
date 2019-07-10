package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
	appName       = "sweatmgr"
)

/* The routing mechanism. Mux helps us define handler functions and the access methods */
func InitRouter() (router *mux.Router) {

	router = mux.NewRouter()

	// No version requirement for /ping
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// Version 1 API management
	v1 := fmt.Sprintf("application/vnd.%s.v1", appName)

	router.HandleFunc("/sweat", createSweatHandler).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/sweat_samples", getSweatSamplesHandler).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/sweat/{id}", getSweatByIdHandler).Methods(http.MethodGet).Headers(versionHeader, v1)

	// Version 2 API management
	v2 := fmt.Sprintf("application/vnd.%s.v2", appName)

	router.HandleFunc("/sweat", createSweatHandler).Methods(http.MethodPost).Headers(versionHeader, v2)
	router.HandleFunc("/sweat_samples", getSweatSamplesHandler).Methods(http.MethodGet).Headers(versionHeader, v2)
	router.HandleFunc("/sweat/{id}", getSweatByIdHandler).Methods(http.MethodGet).Headers(versionHeader, v2)
	router.HandleFunc("/users/{user_id}/sweat", getSweatByUserIdHandler).Methods(http.MethodGet).Headers(versionHeader, v2) // New API
	return
}
