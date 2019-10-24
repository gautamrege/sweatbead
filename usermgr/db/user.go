package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gautamrege/packt/sweatbead/usermgr/logger"
)

const USERS_TABLE string = "users"

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"UserID"`
	Name   string             `bson:"name,omitempty" json:"Name"`
	Device string             `bson:"device,omitempty" json:"Device"`
}

func (mds *MongoDBStorer) Create(ctx context.Context, u User) (err error) {
	_, err = mds.DB.Collection(USERS_TABLE).InsertOne(ctx, u)
	if err != nil {
		logger.Get().Errorf("Error inserting user: %v", err)
		return
	}

	logger.Get().Info("Inserted user into collection")
	return
}

func (mds *MongoDBStorer) ByID(ctx context.Context, id string) (user User, err error) {
	userID, err := primitive.ObjectIDFromHex(id)
	err = mds.DB.Collection(USERS_TABLE).FindOne(ctx, bson.D{{"_id", userID}}).Decode(&user)
	if err != nil {
		logger.Get().Error("User not found in DB")
		return
	}

	return
}

func (mds *MongoDBStorer) List(ctx context.Context) (users []User, err error) {
	collection := mds.DB.Collection(USERS_TABLE)
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
