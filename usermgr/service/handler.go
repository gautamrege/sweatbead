package service

import (
	"net/http"

	"github.com/gautamrege/packt/sweatbead/usermgr/db"
	"github.com/gautamrege/packt/sweatbead/usermgr/logger"
)

// @grpc - enable when required
type GrpcServer struct {
	DB db.Storer
}

// Sample Handler Func template
func pingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		logger.Get().Info("Received ping")
		respBytes := []byte(`{ "Message": "pong" }`)

		// Access the database if required
		// something, err := deps.DB.List(req.Context())

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}
