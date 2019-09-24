package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (mds *MockDB) ListUserSweat(ctx context.Context) (sweats []Sweat, err error) {
	args := mds.Called(ctx)
	return args.Get(0).([]Sweat), args.Error(1)
}

func (mds *MockDB) Create(ctx context.Context, s Sweat) error {
	args := mds.Called(ctx)
	return args.Error(0)
}

func (mds *MockDB) Delete(ctx context.Context, id string) error {
	args := mds.Called(ctx)
	return args.Error(0)
}

func (mds *MockDB) ListAllSweat(ctx context.Context) ([]Sweat, error) {
	args := mds.Called(ctx)
	return args.Get(0).([]Sweat), args.Error(1)
}
