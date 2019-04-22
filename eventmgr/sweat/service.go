package sweat

import (
	"context"

	"github.com/gautamrege/sweatbead/eventmgr/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service interface {
	create(ctx context.Context, req createRequest) (response createUpdateResponse, err error)
	list(ctx context.Context) (response listResponse, err error)
	findByID(ctx context.Context, sweatID primitive.ObjectID) (response findByIDResponse, err error)
}

type sweatService struct {
	store      db.Storer
	logger     *zap.SugaredLogger
	collection *mongo.Collection
}

func (cs *sweatService) create(ctx context.Context, c createRequest) (response createUpdateResponse, err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for sweat create", "msg", err.Error(), "sweat", c)
		return
	}

	sweat, err := cs.store.CreateSweat(ctx, &db.Sweat{
		UserId:      c.UserId,
		Volume:      c.Volume,
		PH:          c.PH,
		Timestamp:   c.Timestamp,
		Moisture:    c.Moisture,
		Temperature: c.Temperature,
	})
	if err != nil {
		cs.logger.Error("Error creating sweat", "err", err.Error())
		return
	}
	response.Sweat = sweat
	response.Message = "Created Successfully!"
	return
}

func (cs *sweatService) list(ctx context.Context) (response listResponse, err error) {
	sweats, err := cs.store.ListSweats(ctx)
	if err != nil {
		cs.logger.Error("Error listing sweats", "err", err.Error())
		return
	}

	response.Sweats = sweats
	return
}

func (cs *sweatService) findByID(ctx context.Context, id primitive.ObjectID) (response findByIDResponse, err error) {
	sweat, err := cs.store.FindSweatByID(ctx, id)
	if err == db.ErrSweatNotExist {
		cs.logger.Error("No sweat present", "err", err.Error())
		return response, errNoSweatId
	}

	if err != nil {
		cs.logger.Error("Error finding sweat", "err", err.Error(), "sweat_id", id)
		return
	}

	response.Sweat = sweat
	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &sweatService{
		store:  s,
		logger: l,
	}
}
