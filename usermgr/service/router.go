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

	root := mux.NewRouter()

	// Make a path prefix that will be assigned to this microservice
	router = root.PathPrefix("/usermgr").Subrouter()

	// Version 1 API management
	v1 := fmt.Sprintf("application/vnd.%s.v1", appName)

	router.HandleFunc("/user", createUserHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1) // create users

	// Note: Get user(s) will be supported only via gRPC

	return
}
