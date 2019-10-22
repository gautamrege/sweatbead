package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
	appName       = "usermgr"
)

/* The routing mechanism. Mux helps us define handler functions and the access methods */
func InitRouter(deps Dependencies) (router *mux.Router) {

	router = mux.NewRouter()

	// Version 1 API management
	v1 := fmt.Sprintf("application/vnd.%s.v1", appName)

	router.HandleFunc("/ping", pingHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)

	// TODO: Add Handlers
	//router.HandleFunc("/users", getUsersHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	//router.HandleFunc("/user", createUserHandlea(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)

	return
}
