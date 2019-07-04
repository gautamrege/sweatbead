package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	db *mongo.Database
)

const (
	user     string = ""
	password string = ""
	host     string = "localhost"
	port     int32  = 27017
	name     string = "sampledb"
)

// Singleton instance method accessible from other packages
func GetDB() (db *mongo.Database, err error) {
	if db != nil {
		return
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authMechanism=SCRAM-SHA-1", user, password, host, port, name)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return
	}

	fmt.Println("Connected To MongoDB")
	db = client.Database(name)
	return
}
