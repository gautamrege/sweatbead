package user

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func makeHTTPCall(handler http.HandlerFunc, method, path, body string) (rr *httptest.ResponseRecorder) {
	request := []byte(body)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(request))
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return
}

// Create:
func TestSuccessfullCreate(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("create", mock.Anything, mock.Anything).Return(nil)

	rr := makeHTTPCall(Create(cs), http.MethodPost, "/users", `{"name":"Sports"}`)

	checkResponseCode(t, http.StatusCreated, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenInvalidRequestBody(t *testing.T) {
	cs := &UserServiceMock{}

	rr := makeHTTPCall(Create(cs), http.MethodPost, "/users", `{"name":"",}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenEmptyName(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("create", mock.Anything, mock.Anything).Return(errEmptyName)

	rr := makeHTTPCall(Create(cs), http.MethodPost, "/users", `{"name":""}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestCreateWhenInternalError(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("create", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

	rr := makeHTTPCall(Create(cs), http.MethodPost, "/users", `{"name":"Sports"}`)

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

// List :
func TestSuccessfullList(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("list", mock.Anything).Return(mock.Anything, nil)

	rr := makeHTTPCall(List(cs), http.MethodGet, "/users", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestListWhenNoUsers(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("list", mock.Anything).Return(mock.Anything, errNoUsers)

	rr := makeHTTPCall(List(cs), http.MethodGet, "/users", "")

	checkResponseCode(t, http.StatusNotFound, rr.Code)
	cs.AssertExpectations(t)
}

func TestListInternalError(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("list", mock.Anything).Return(mock.Anything, errors.New("Internal Error"))

	rr := makeHTTPCall(List(cs), http.MethodGet, "/users", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

//FindById
//not bad reqe
//not find err
func TestSuccessfullFindByID(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("findByID", mock.Anything, mock.Anything).Return(mock.Anything, nil)

	rr := makeHTTPCall(FindByID(cs), http.MethodGet, "/users/1", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestFindByIDWhenIDNotExist(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("findByID", mock.Anything, mock.Anything).Return(mock.Anything, errNoUserId)

	rr := makeHTTPCall(FindByID(cs), http.MethodGet, "/users/1", "")

	checkResponseCode(t, http.StatusNotFound, rr.Code)
	cs.AssertExpectations(t)
}

func TestFindByIdWhenInternalError(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("findByID", mock.Anything, mock.Anything).Return(mock.Anything, errors.New("Internal Error"))

	rr := makeHTTPCall(FindByID(cs), http.MethodGet, "/users/1", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

//DeleteByID
func TestSuccessfullDeleteByID(t *testing.T) {
	cs := &UserServiceMock{}

	cs.On("deleteByID", mock.Anything, mock.Anything).Return(nil)

	rr := makeHTTPCall(DeleteByID(cs), http.MethodDelete, "/users/1", "")

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestDeleteByIDWhenIDNotExist(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("deleteByID", mock.Anything, mock.Anything).Return(errNoUserId)

	rr := makeHTTPCall(DeleteByID(cs), http.MethodDelete, "/users/1", "")

	checkResponseCode(t, http.StatusNotFound, rr.Code)
	cs.AssertExpectations(t)
}

func TestDeleteByIDWhenInternalError(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("deleteByID", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

	rr := makeHTTPCall(DeleteByID(cs), http.MethodDelete, "/users/1", "")

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}

func TestSuccessfullUpdate(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("update", mock.Anything, mock.Anything).Return(nil)

	rr := makeHTTPCall(Update(cs), http.MethodPut, "/users", `{"id":"1", "name":"sports"}`)

	checkResponseCode(t, http.StatusOK, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenInvalidRequestBody(t *testing.T) {
	cs := &UserServiceMock{}

	rr := makeHTTPCall(Update(cs), http.MethodPut, "/users", `{"id":"1", "name":"sports",}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenEmptyID(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("update", mock.Anything, mock.Anything).Return(errEmptyID)

	rr := makeHTTPCall(Update(cs), http.MethodPut, "/users", `{"name":"Sports"}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenEmptyName(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("update", mock.Anything, mock.Anything).Return(errEmptyName)

	rr := makeHTTPCall(Update(cs), http.MethodPut, "/users", `{"id":"1"}`)

	checkResponseCode(t, http.StatusBadRequest, rr.Code)
	cs.AssertExpectations(t)
}

func TestUpdateWhenInternalError(t *testing.T) {
	cs := &UserServiceMock{}
	cs.On("update", mock.Anything, mock.Anything).Return(errors.New("Internal Error"))

	rr := makeHTTPCall(Update(cs), http.MethodPut, "/users", `{"id":"1"}`)

	checkResponseCode(t, http.StatusInternalServerError, rr.Code)
	cs.AssertExpectations(t)
}
