package app

import (
	"context"
	"fmt"
	"time"

	"github.com/gautamrege/sweatbead/eventmgr/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

var (
	db     *mongo.Database
	client *mongo.Client
	logger *zap.SugaredLogger
	ctx    context.Context
)

func Init() {
	InitLogger()

	err := initDB()
	if err != nil {
		panic(err)
	}
}

func InitLogger() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	logger = zapLogger.Sugar()
}

func initDB() (err error) {
	dbConfig := config.Database()
	client, err := mongo.NewClient(options.Client().ApplyURI(dbConfig.ConnectionURL()))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		return err
	}

	fmt.Println("Connected To MongoDB")
	db = client.Database(dbConfig.DbName())
	return
}

func GetLogger() *zap.SugaredLogger {
	return logger
}

func GetDB() *mongo.Database {
	return db
}

func GetCollection(name string) *mongo.Collection {
	collection := db.Collection(name)
	return collection
}

func Close() {
	logger.Sync()
	//db.Close()
}
