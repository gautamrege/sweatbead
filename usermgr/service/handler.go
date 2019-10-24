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

// Create a user
func createUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

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
	})
}
