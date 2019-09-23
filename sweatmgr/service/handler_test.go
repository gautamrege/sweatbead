package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/config"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	testUser db.User
)

func init() {
	config.Load()
	logger.Init()
	db.Init()
}

func TestCreateUser(t *testing.T) {
	var jsonStr = []byte(`{ "Name": "Test User", "Device": "Fitbit test" }`)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createUserHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestFetchUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsersHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var users []db.User

	err = json.Unmarshal([]byte(rr.Body.String()), &users)
	if err != nil {
		t.Fatal(err)
	}

	if len(users) < 1 {
		t.Errorf("No sweat samples. Expected at least 1")
	}

	testUser = users[len(users)-1]
}

func TestCreateSweat(t *testing.T) {
	var jsonStr = []byte(`{ "glucose": 1.12, "sodium": 0.98, "chloride": 0.003 }`)

	req, err := http.NewRequest("POST", "/sweat", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("UserID", testUser.ID.Hex())
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

func TestUserSweat(t *testing.T) {
	mockDB := db.MockDB{}
	deps := Dependencies{
		DB: &mockDB,
	}

	dummy := []db.Sweat{db.Sweat{Glucose: 0.92}}
	mockDB.On("ListUserSweat", mock.Anything).Return(dummy, nil)

	req, err := http.NewRequest("GET", "/user/sweat", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("UserID", testUser.ID.Hex())
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getSweatByUserIdHandler(deps))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var sweats []db.Sweat
	err = json.Unmarshal([]byte(rr.Body.String()), &sweats)
	assert.Nil(t, err)

	assert.ElementsMatch(t, sweats, dummy)
	mockDB.AssertExpectations(t)
}

func TestUserSweatInvalid(t *testing.T) {
	mockDB := db.MockDB{}
	deps := Dependencies{
		DB: &mockDB,
	}

	dummy := []db.Sweat{}
	mockDB.On("ListUserSweat", mock.Anything).Return(dummy, errors.New("No DB"))

	req, err := http.NewRequest("GET", "/user/sweat", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("UserID", testUser.ID.Hex())
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getSweatByUserIdHandler(deps))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// ensure that all the mock methods are called
	mockDB.AssertExpectations(t)
}
