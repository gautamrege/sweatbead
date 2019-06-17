package service

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Message string `json:"message"`
}

/* Each handler will then respond with marshalled Json data */
func pingHandler(rw http.ResponseWriter, req *http.Request) {
	response := PingResponse{Message: "pong"}

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
