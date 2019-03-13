package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type StorerMock struct {
	mock.Mock
}

func (m *StorerMock) CreateUser(ctx context.Context, user *User) (err error) {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *StorerMock) ListUsers(ctx context.Context) (users []User, err error) {
	args := m.Called(ctx)
	users, _ = args.Get(0).([]User)
	return users, args.Error(1)
}

func (m *StorerMock) FindUserByID(ctx context.Context, id string) (user User, err error) {
	args := m.Called(ctx, id)
	user, _ = args.Get(0).(User)
	return user, args.Error(1)
}

func (m *StorerMock) DeleteUserByID(ctx context.Context, id string) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *StorerMock) UpdateUser(ctx context.Context, user *User) (err error) {
	args := m.Called(ctx, user)
	return args.Error(0)
}
