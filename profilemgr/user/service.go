package user

import (
	"context"

	"github.com/gautamrege/sweatbead/profilemgr/db"
	"go.uber.org/zap"
)

type Service interface {
	list(ctx context.Context) (response listResponse, err error)
	create(ctx context.Context, req createRequest) (err error)
	findByID(ctx context.Context, id string) (response findByIDResponse, err error)
	deleteByID(ctx context.Context, id string) (err error)
	update(ctx context.Context, req updateRequest) (err error)
}

type userService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *userService) list(ctx context.Context) (response listResponse, err error) {
	users, err := cs.store.ListUsers(ctx)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUsers
	}
	if err != nil {
		cs.logger.Error("Error listing users", "err", err.Error())
		return
	}

	response.Users = users
	return
}

func (cs *userService) create(ctx context.Context, c createRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for user create", "msg", err.Error(), "user", c)
		return
	}

	err = cs.store.CreateUser(ctx, &db.User{
		Name: c.Name,
	})
	if err != nil {
		cs.logger.Error("Error creating user", "err", err.Error())
		return
	}
	return
}

func (cs *userService) update(ctx context.Context, c updateRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Error("Invalid Request for user update", "err", err.Error(), "user", c)
		return
	}

	err = cs.store.UpdateUser(ctx, &db.User{
		ID:   c.ID,
		Name: c.Name,
	})
	if err != nil {
		cs.logger.Error("Error updating user", "err", err.Error(), "user", c)
		return
	}

	return
}

func (cs *userService) findByID(ctx context.Context, id string) (response findByIDResponse, err error) {
	user, err := cs.store.FindUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "user_id", id)
		return
	}

	response.User = user
	return
}

func (cs *userService) deleteByID(ctx context.Context, id string) (err error) {
	err = cs.store.DeleteUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("User Not present", "err", err.Error(), "user_id", id)
		return errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error deleting user", "err", err.Error(), "user_id", id)
		return
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userService{
		store:  s,
		logger: l,
	}
}
