package db

import (
	"context"
	"time"

	"github.com/gautamrege/sweatbead/eventmgr/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sweat struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // Here omitempty is must otherwise Go will set zero value as ID to document
	UserId      string             `bson:"user_id"`       // Storing UserID as string as its responsibility of 3rd party apis to typecast as and when needed
	Volume      float32            `bson:"volume"`
	PH          float32            `bson:"pH"`
	Timestamp   int64              `bson:"timestamp"`
	Moisture    float32            `bson:"moisture"`
	Temperature float32            `bson:"temperature"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

const (
	sweatCollectionName = `sweats`
)

func (s *store) CreateSweat(ctx context.Context, sweat *Sweat) (sweatInfo Sweat, err error) {
	now := time.Now()
	sweat.CreatedAt = now
	sweat.UpdatedAt = now
	collection := app.GetCollection(sweatCollectionName)
	res, err := collection.InsertOne(ctx, sweat)
	if err != nil {
		return
	}
	id := res.InsertedID
	err = collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&sweatInfo)
	return // Here we are not handling err explicilty as we are returning in success & err case. In future if we add some other logic before successful return we will need to handle err as well
}

func (s *store) ListSweats(ctx context.Context) (sweats []Sweat, err error) {
	collection := app.GetCollection(sweatCollectionName)
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			return err
		}
		defer cur.Close(ctx)
		var elem Sweat
		for cur.Next(ctx) {
			err = cur.Decode(&elem)
			sweats = append(sweats, elem)
		}
		if err := cur.Err(); err != nil {
		}
		return err
	})
	return
}
