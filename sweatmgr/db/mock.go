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
