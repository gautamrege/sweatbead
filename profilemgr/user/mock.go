package user

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) create(ctx context.Context, c createRequest) (err error) {
	args := m.Called(ctx, c)
	return args.Error(0)
}

func (m *UserServiceMock) list(ctx context.Context) (response listResponse, err error) {
	args := m.Called(ctx)
	response, _ = args.Get(0).(listResponse)
	return response, args.Error(1)
}

func (m *UserServiceMock) update(ctx context.Context, c updateRequest) (err error) {
	args := m.Called(ctx, c)
	return args.Error(0)
}

func (m *UserServiceMock) findByID(ctx context.Context, id string) (response findByIDResponse, err error) {
	args := m.Called(ctx)
	response, _ = args.Get(0).(findByIDResponse)
	return response, args.Error(1)
}

func (m *UserServiceMock) deleteByID(ctx context.Context, id string) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}
