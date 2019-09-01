package db

import (
	"context"
	"errors"

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

func (s *Sweat) Create(ctx context.Context) (err error) {
	userid := ""
	if ctx.Value("UserID") != nil { // verify it exists
		userid = ctx.Value("UserID").(string)
	}

	if userid == "" {
		logger.Get().Error("User not specified in context")
		return errors.New("User not found")
	}

	objID, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		logger.Get().Error("UserID is invalid")
		return errors.New("User not found")
	}

	var user User
	err = db.Collection(USERS_TABLE).FindOne(ctx, bson.D{{"_id", objID}}).Decode(&user)
	if err != nil {
		logger.Get().Error("User not found in DB")
		return err
	}

	s.UserID = user.ID
	s.CreatedAt = time.Now()
	collection := db.Collection(SWEAT_TABLE)
	_, err = collection.InsertOne(ctx, s)
	if err != nil {
		logger.Get().Infof("Error inserting sweat: %v", s)
		return
	}

	logger.Get().Info("Inserted sweat into collection")
	return
}

func (s *Sweat) Delete() (err error) {
	return
}

func ListAllSweat() (sweats []Sweat, err error) {
	collection := db.Collection(SWEAT_TABLE)
	ctx := context.TODO()
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

func GetSweatByID(id string) (s Sweat, err error) {
	return
}

func GetSweatByTime(start, end time.Time) (s []Sweat, err error) {
	return
}
