package server

import (
	"fmt"
	"net/http"

	"github.com/gautamrege/sweatbead/eventmgr/api"
	"github.com/gautamrege/sweatbead/eventmgr/config"
	"github.com/gautamrege/sweatbead/eventmgr/sweat"
	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())
	fmt.Println(v1)

	// TODO: add doc
	// v2 := fmt.Sprintf("application/vnd.%s.v2", config.AppName())

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// Event
	router.HandleFunc("/sweats", sweat.Create(dep.SweatService)).Methods(http.MethodPost).Headers(versionHeader, v1)
	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
