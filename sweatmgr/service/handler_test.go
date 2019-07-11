package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
)

func TestCreateSweat(t *testing.T) {
	var jsonStr = []byte(`{ "glucose": 1.12, "sodium": 0.98, "chloride": 0.003 }`)

	req, err := http.NewRequest("POST", "/entry", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createSweatHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSweatSamples(t *testing.T) {
	req, err := http.NewRequest("GET", "/sweat_samples", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getSweatSamplesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var sweats []db.Sweat

	err = json.Unmarshal([]byte(rr.Body.String()), &sweats)
	if err != nil {
		t.Fatal(err)
	}

	if len(sweats) < 1 {
		t.Errorf("No sweat samples. Expected at least 1")
	}
}
