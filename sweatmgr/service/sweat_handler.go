package service

import (
	"encoding/json"
	"net/http"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
)

func createSweatHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)

		var s db.Sweat
		err := decoder.Decode(&s)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Verify this user exists - fetch user from UserMgr service via gRPC
		_, ok := deps.UserMgr.GetUser(req.Header.Get("UserID"))
		if ok != nil {
			logger.Get().Error("User not found")
			rw.WriteHeader(http.StatusForbidden)
			return
		}

		req = WithUserContext(req)
		err = deps.DB.Create(req.Context(), s)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
	})
}

func getSweatSamplesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		sweats, err := deps.DB.ListAllSweat(req.Context())
		if err != nil {
			logger.Get().Info("Error fetching data", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(sweats)
		if err != nil {
			logger.Get().Info("Error marshaling data", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

// In this function, we need to call the database and for a scalable microservice,
// we will also need to call other microservices! So, we need to have handle multiple
// dependencies
func getSweatByUserIdHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		req = WithUserContext(req)

		sweats, err := deps.DB.ListUserSweat(req.Context())
		if err != nil {
			logger.Get().Info("Error fetching data", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(sweats)
		if err != nil {
			logger.Get().Info("Error marshaling data", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}
