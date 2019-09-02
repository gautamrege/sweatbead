package service

import (
	"encoding/json"
	"net/http"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
)

// @Title createUser
// @Description create User
// @Accept  json
// @Success 200 {object}
// @Failure 400 {object}
// @Router /user [post]
func createUserHandler(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var s db.User
	err := decoder.Decode(&s)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.Create(req.Context())
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}

func getUsersHandler(rw http.ResponseWriter, req *http.Request) {
	users, err := db.ListAllUsers(req.Context())
	if err != nil {
		logger.Get().Info("Error fetching data", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBytes, err := json.Marshal(users)
	if err != nil {
		logger.Get().Info("Error marshaling data", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(respBytes)

}
