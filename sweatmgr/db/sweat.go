package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
)

const SWEAT_TABLE string = "sweat" // the collection name

type Sweat struct {
	// Database specific fields.
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userid,omitempty" json:"UserID"`
	CreatedAt time.Time          `bson:"created_at"`

	// Potential disease Diagnosis
	Glucose  float32 `bson:"glucose" json:"Glucose,omitempty"`   // excess indicates diabetes
	Chloride float32 `bson:"chloride" json:"Chloride,omitempty"` // excess indicates cystic fibrosis

	// Electrolytes
	Sodium    float32 `bson:"sodium" json:"Sodium,omitempty"`
	Potassium float32 `bson:"potassium" json:"Potassium,omitempty"` // excess indicates exercise / workout
	Magnesium float32 `bson:"magnesium" json:"Magnesium,omitempty"` // excess indicates exercise / workout
	Calcium   float32 `bson:"calcium" json:"Calcium,omitempty"`     // excess indicates exercise / workout

	// Environmental conditions and determining criteria
	Humidity         float32 `bson:"humidity" json:"Humidity,omitempty"`                // high humidity increseas sweating
	RoomTemperatureF float32 `bson:"room_temperature" json:"RoomTemperature,omitempty"` // cooler room temperature with sweat indicates hyperdidrosis
	BodyTemperatureF float32 `bson:"body_temperature" json:"BodyTemperature,omitempty"` // high body temperature with sweat indicates fever
	HeartBeat        int32   `bson:"heartbeat" json:"HeartBeat,omitempty"`              // sweating without apparent reason is an alarming condition!
}

func (mds *MongoDBStorer) Delete(ctx context.Context, id string) (err error) {
	return
}

func (mds *MongoDBStorer) ListAllSweat(ctx context.Context) (sweats []Sweat, err error) {
	collection := mds.DB.Collection(SWEAT_TABLE)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	var elem Sweat
	for cur.Next(ctx) {
		err = cur.Decode(&elem)
		sweats = append(sweats, elem)
	}
	if err = cur.Err(); err != nil {
		logger.Get().Infof("Error in listing data: ", err)
		return
	}
	return
}

func (mds *MongoDBStorer) Create(ctx context.Context, s Sweat) (err error) {
	s.UserID = userIDFromContext(ctx)
	s.CreatedAt = time.Now()
	collection := mds.DB.Collection(SWEAT_TABLE)
	_, err = collection.InsertOne(ctx, s)
	if err != nil {
		logger.Get().Infof("Error inserting sweat: %v", s)
		return
	}

	logger.Get().Info("Inserted sweat into collection")
	return
}

func (mds *MongoDBStorer) ListUserSweat(ctx context.Context) (sweats []Sweat, err error) {
	user, err := GetUserByID(ctx, userIDFromContext(ctx))
	if err != nil {
		return
	}

	filter := bson.D{
		{"userid", user.ID},
	}
	cur, err := mds.DB.Collection(SWEAT_TABLE).Find(ctx, filter)
	if err != nil {
		return
	}
	defer cur.Close(ctx)

	var elem Sweat
	for cur.Next(ctx) {
		err = cur.Decode(&elem)
		sweats = append(sweats, elem)
	}
	if err = cur.Err(); err != nil {
		logger.Get().Infof("Error in listing data: ", err)
		return
	}
	return
}
