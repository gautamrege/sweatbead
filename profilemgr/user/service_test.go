package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/gautamrege/sweatbead/profilemgr/app"
	"github.com/gautamrege/sweatbead/profilemgr/db"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfullCreateService(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var tests = []struct {
		contx    context.Context
		req      createRequest
		expected error
	}{
		{ctx, createRequest{Name: "Sports"}, nil},
		{ctx, createRequest{Name: "Reading"}, nil},
	}
	for _, test := range tests {
		sm.On("CreateUser", test.contx, mock.Anything).Return(nil)
		assert.Equal(cs.create(test.contx, test.req), test.expected)
		sm.AssertExpectations(t)
	}
}

func TestCreateServiceWhenEmptyName(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      createRequest
		expected error
	}{
		ctx,
		createRequest{Name: ""},
		errEmptyName,
	}

	assert.Equal(cs.create(test.contx, test.req), test.expected)
	sm.AssertExpectations(t)
}

func TestCreateServiceWhenInternalError(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      createRequest
		expected error
	}{
		ctx,
		createRequest{Name: "Sports"},
		errors.New("Internal Error"),
	}

	sm.On("CreateUser", test.contx, mock.Anything).Return(errors.New("Internal Error"))
	assert.Equal(cs.create(test.contx, test.req), test.expected)
	sm.AssertExpectations(t)
}

func TestSuccessfullListService(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		expected error
	}{ctx, nil}
	sm.On("ListUsers", test.contx).Return(mock.Anything, nil)
	_, err := cs.list(test.contx)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestListServiceWhenUsersNotExists(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		expected error
	}{ctx, errNoUsers}
	sm.On("ListUsers", test.contx).Return(mock.Anything, db.ErrUserNotExist)
	_, err := cs.list(test.contx)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestListServiceWhenInternalError(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		expected error
	}{ctx, errors.New("Internal Error")}
	sm.On("ListUsers", test.contx).Return(mock.Anything, errors.New("Internal Error"))
	_, err := cs.list(test.contx)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestSuccessfullUpdateService(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      updateRequest
		expected error
	}{ctx, updateRequest{ID: "1", Name: "Sports"}, nil}
	sm.On("UpdateUser", test.contx, mock.Anything).Return(nil)
	assert.Equal(cs.update(test.contx, test.req), test.expected)
	sm.AssertExpectations(t)
}

func TestUpdateServiceWhenEmptyID(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      updateRequest
		expected error
	}{ctx, updateRequest{ID: "", Name: "Sports"}, errEmptyID}
	assert.Equal(cs.update(test.contx, test.req), test.expected)
	sm.AssertExpectations(t)
}

func TestUpdateServiceWhenEmptyName(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      updateRequest
		expected error
	}{ctx, updateRequest{ID: "1", Name: ""}, errEmptyName}
	assert.Equal(cs.update(test.contx, test.req), test.expected)
	sm.AssertExpectations(t)
}

func TestUpdateServiceWhenInternalError(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		req      updateRequest
		expected error
	}{ctx, updateRequest{ID: "1", Name: "Sports"}, errors.New("Internal Error")}
	sm.On("UpdateUser", test.contx, mock.Anything).Return(errors.New("Internal Error"))
	assert.Equal(cs.update(test.contx, test.req), test.expected)
	sm.AssertExpectations(t)
}

func TestSuccessfullFindByIDService(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", nil}
	sm.On("FindUserByID", test.contx, test.id).Return(mock.Anything, nil)
	_, err := cs.findByID(test.contx, test.id)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestFindByIDServiceWhenUserNotExist(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", errNoUserId}
	sm.On("FindUserByID", test.contx, mock.Anything).Return(mock.Anything, db.ErrUserNotExist)
	_, err := cs.findByID(test.contx, test.id)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestFindByIDServiceWhenInternalError(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", errors.New("Internal Error")}
	sm.On("FindUserByID", test.contx, mock.Anything).Return(mock.Anything, errors.New("Internal Error"))
	_, err := cs.findByID(test.contx, test.id)
	assert.Equal(err, test.expected)
	sm.AssertExpectations(t)
}

func TestSuccessfullDeleteByIDService(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", nil}
	sm.On("DeleteUserByID", test.contx, test.id).Return(nil)
	assert.Equal(cs.deleteByID(test.contx, test.id), test.expected)
	sm.AssertExpectations(t)
}

func TestDeleteByIDServiceWhenUserNotExist(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", errNoUserId}
	sm.On("DeleteUserByID", test.contx, test.id).Return(db.ErrUserNotExist)
	assert.Equal(cs.deleteByID(test.contx, test.id), test.expected)
	sm.AssertExpectations(t)
}

func TestDeleteByIDServiceWhenInternalError(t *testing.T) {
	app.InitLogger()
	sm := &db.StorerMock{}
	l := app.GetLogger()
	cs := NewService(sm, l)

	ctx := context.Background()
	assert := assert.New(t)

	var test = struct {
		contx    context.Context
		id       string
		expected error
	}{ctx, "1", errors.New("Internal Error")}
	sm.On("DeleteUserByID", test.contx, test.id).Return(errors.New("Internal Error"))
	assert.Equal(cs.deleteByID(test.contx, test.id), test.expected)
	sm.AssertExpectations(t)
}
