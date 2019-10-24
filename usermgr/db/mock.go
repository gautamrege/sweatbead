package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (mds *MockDB) Create(ctx context.Context, u User) error {
	args := mds.Called(ctx)
	return args.Error(0)

}
func (mds *MockDB) ByID(ctx context.Context, id string) (User, error) {
	args := mds.Called(ctx)
	return args.Get(0).(User), args.Error(1)
}

func (mds *MockDB) List(ctx context.Context) ([]User, error) {
	args := mds.Called(ctx)
	return args.Get(0).([]User), args.Error(1)
}
