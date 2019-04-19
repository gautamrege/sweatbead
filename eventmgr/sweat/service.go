package sweat

import (
	"context"

	"github.com/gautamrege/sweatbead/eventmgr/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Service interface {
	create(ctx context.Context, req createRequest) (err error)
}

type sweatService struct {
	store      db.Storer
	logger     *zap.SugaredLogger
	collection *mongo.Collection
}

func (cs *sweatService) create(ctx context.Context, c createRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for sweat create", "msg", err.Error(), "sweat", c)
		return
	}

	err = cs.store.CreateSweat(ctx, &db.Sweat{
		UserId:      c.UserId,
		Volume:      c.Volume,
		PH:          c.PH,
		Timestamp:   c.Timestamp,
		Moisture:    c.Moisture,
		Temperature: c.Temperature,
	})
	if err != nil {
		cs.logger.Error("Error creating user", "err", err.Error())
		return
	}
	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &sweatService{
		store:  s,
		logger: l,
	}
}
