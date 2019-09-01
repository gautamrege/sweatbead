package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
)

const USERS_TABLE string = "users"

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"UserID"`
	Name   string             `bson:"name,omitempty" json:"Name"`
	Device string             `bson:"device,omitempty" json:"Device"`
}

func (u *User) Create(ctx context.Context) (err error) {
	_, err = db.Collection(USERS_TABLE).InsertOne(ctx, u)
	if err != nil {
		logger.Get().Errorf("Error inserting user: %v", err)
		return
	}

	logger.Get().Info("Inserted user into collection")
	return
}

func ListAllUsers(ctx context.Context) (users []User, err error) {
	collection := db.Collection(USERS_TABLE)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	var elem User
	for cur.Next(ctx) {
		err = cur.Decode(&elem)
		users = append(users, elem)
	}
	if err = cur.Err(); err != nil {
		logger.Get().Infof("Error in listing data: ", err)
		return
	}
	return

}
