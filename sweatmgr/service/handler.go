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

	// Testing: DB operation - start
	s := db.Sweat{
		Glucose:          0.01,
		Chloride:         0.002,
		Sodium:           0.9,
		HeartBeat:        72,
		RoomTemperatureF: 76,
	}

	_ = s.Create()
	// Testing: DB operation - end

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
