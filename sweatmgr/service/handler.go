package service

import (
	"context"
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
		status = http.StatusInternalServerError
		panic(err)
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func WithUserContext(req *http.Request) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "UserID", req.Header.Get("UserID"))
	return req.WithContext(ctx)
}
