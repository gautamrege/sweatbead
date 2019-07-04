package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"
)

const SWEAT_TABLE string = "sweat" // the collection name

type Sweat struct {
	// Database specific fields.
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userid,omitempty" json:"userid"`
	CreatedAt time.Time          `bson:"created_at"`

	// Potential disease Diagnosis
	Glucose  float32 `bson:"glucose" json:"glucose,omitempty"`   // excess indicates diabetes
	Chloride float32 `bson:"chloride" json:"chloride,omitempty"` // excess indicates cystic fibrosis

	// Electrolytes
	Sodium    float32 `bson:"sodium" json:"sodium,omitempty"`
	Potassium float32 `bson:"potassium" json:"potassium,omitempty"` // excess indicates exercise / workout
	Magnesium float32 `bson:"magnesium" json:"magnesium,omitempty"` // excess indicates exercise / workout
	Calcium   float32 `bson:"calcium" json:"calcium,omitempty"`     // excess indicates exercise / workout

	// Environmental conditions and determining criteria
	Humidity        float32 `bson:"humidity" json:"humidity,omitempty"`                 // high humidity increseas sweating
	RoomTemperature float32 `bson:"room_temperature" json:"room_temperature,omitempty"` // cooler room temperature with sweat indicates hyperdidrosis
	BodyTemperature float32 `bson:"body_temperature" json:"body_temperature,omitempty"` // high body temperature with sweat indicates fever
	HeartBeat       int32   `bson:"heartbeat" json:"heartbeat,omitempty"`               // sweating without apparent reason is an alarming condition!
}

func (s *Sweat) Create() (err error) {
	db, err := GetDB()
	if err != nil {
		fmt.Println("No Database connection")
		return err
	}

	s.CreatedAt = time.Now()
	collection := db.Collection(SWEAT_TABLE)
	_, err = collection.InsertOne(context.TODO(), s)
	if err != nil {
		fmt.Printf("Error inserting sweat: %v", s)
		return
	}

	fmt.Println("Inserted sweat into collection")
	return
}

func (s *Sweat) Delete() (err error) {
	return
}

func GetSweatByID(id string) (s Sweat, err error) {
	return
}

func GetSweatByTime(start, end time.Time) (s []Sweat, err error) {
	return
}
