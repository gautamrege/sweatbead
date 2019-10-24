package service

import (
	"encoding/json"
	"net/http"

	"github.com/gautamrege/packt/sweatbead/usermgr/db"
)

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

		err = deps.DB.Create(req.Context(), s)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
	})
}
