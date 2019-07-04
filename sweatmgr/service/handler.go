package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
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
		status = http.StatusInternalServerError
		panic(err)
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func createSweatHandler(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	status := http.StatusOK

	var s db.Sweat

	err := decoder.Decode(&s)
	if err != nil {
		status = http.StatusInternalServerError
		panic(err)
	}
	fmt.Println(s)
	err = s.Create()
	if err != nil {
		status = http.StatusInternalServerError
		panic(err)
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
}
