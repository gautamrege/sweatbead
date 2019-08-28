package service

import (
	"encoding/json"
	"net/http"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
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
	logger.Get().Info(s)
	err = s.Create()
	if err != nil {
		status = http.StatusInternalServerError
		panic(err)
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
}

func getSweatSamplesHandler(rw http.ResponseWriter, req *http.Request) {
	status := http.StatusOK
	sweats, err := db.ListAllSweat()
	if err != nil {
		logger.Get().Info("Error fetching data", err)
		status = http.StatusInternalServerError
	}

	respBytes, err := json.Marshal(sweats)
	if err != nil {
		logger.Get().Info("Error marshaling data", err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)

}

func getSweatByIdHandler(rw http.ResponseWriter, req *http.Request) {
}

func getSweatByUserIdHandler(rw http.ResponseWriter, req *http.Request) {
}
