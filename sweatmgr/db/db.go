package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/config"
)

var (
	db *mongo.Database
)

// Singleton instance method accessible from other packages
func GetDB() (db *mongo.Database, err error) {
	if db != nil {
		return
	}

	user := config.ReadEnvString("DB_USER")
	host := config.ReadEnvString("DB_HOST")
	port := config.ReadEnvInt("DB_PORT")
	name := config.ReadEnvString("DB_NAME")
	password := config.ReadEnvString("DB_PASSWORD")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authMechanism=%s", user, password, host, port, name, config.ReadEnvString("DB_AUTH_MECH"))

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
