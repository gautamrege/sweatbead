package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/config"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
)

var (
	db *mongo.Database
)

type Storer interface {
	Create(context.Context, Sweat) error
	Delete(context.Context, string) error
	ListAllSweat(context.Context) ([]Sweat, error)
	ListUserSweat(context.Context) ([]Sweat, error)
}

type MongoDBStorer struct {
	DB *mongo.Database
}

// Singleton instance method accessible from other packages
func Init() {
	user := config.ReadEnvString("DB_USER")
	host := config.ReadEnvString("DB_HOST")
	port := config.ReadEnvInt("DB_PORT")
	name := config.ReadEnvString("DB_NAME")
	password := config.ReadEnvString("DB_PASSWORD")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authMechanism=%s", user, password, host, port, name, config.ReadEnvString("DB_AUTH_MECH"))

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Get().Fatal("Cannot initialize database")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logger.Get().Fatal("Cannot initialize database context")
		return
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		logger.Get().Fatal("Cannot connect to database")
		return
	}

	logger.Get().Info("Connected To MongoDB")
	db = client.Database(name)
	return
}

func Get() *mongo.Database {
	if db == nil {
		logger.Get().Fatal("Database not initialized")
		return nil
	}

	return db
}

func GetStorer(db *mongo.Database) Storer {
	return &MongoDBStorer{db}
}

func userIDFromContext(ctx context.Context) (userID primitive.ObjectID) {
	userid := ""
	if ctx.Value("UserID") != nil { // verify it exists
		userid = ctx.Value("UserID").(string)
	}

	if userid == "" {
		logger.Get().Error("User not specified in context")
		return
	}

	userID, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		logger.Get().Errorf("UserID is invalid: %v", err)
		return
	}

	return
}
