package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sweat struct {
	Id          primitive.ObjectID `bson:_id`
	Volume      float32            `bson:volume`
	PH          float32            `bson:pH`
	Timestamp   int64              `bson:timestamp`
	Moisture    float32            `bson:moisture`
	Temperature float32            `bson:temperature`
}
