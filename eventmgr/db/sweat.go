package db

import (
	"context"
	"time"

	"github.com/gautamrege/sweatbead/eventmgr/app"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sweat struct {
	Id          primitive.ObjectID `bson:_id`
	UserId      int64              `bson: user_id`
	Volume      float32            `bson:volume`
	PH          float32            `bson:pH`
	Timestamp   int64              `bson:timestamp`
	Moisture    float32            `bson:moisture`
	Temperature float32            `bson:temperature`
	CreatedAt   time.Time          `db:"created_at"`
	UpdatedAt   time.Time          `db:"updated_at"`
}

const (
	collectionName = `sweats`
)

func (s *store) CreateSweat(ctx context.Context, sweat *Sweat) (err error) {
	now := time.Now()
	sweat.CreatedAt = now
	sweat.UpdatedAt = now
	collection := app.GetCollection(collectionName)
	_, err = collection.InsertOne(ctx, sweat)
	if err != nil {
		return
	}
	//id := res.InsertedID
	return
}
