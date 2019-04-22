package sweat

import (
	"github.com/gautamrege/sweatbead/eventmgr/db"
)

type createRequest struct {
	UserId      string  `json: user_id`
	Volume      float32 `json:volume`
	PH          float32 `json:pH`
	Timestamp   int64   `json:timestamp`
	Moisture    float32 `json:moisture`
	Temperature float32 `json:temperature`
}

type createUpdateResponse struct {
	Sweat   db.Sweat
	Message string
}

func (cr createRequest) Validate() (err error) {
	if cr.UserId == "" {
		return errEmptyUserID
	}
	if cr.Volume == 0.0 {
		return errEmptyVolume
	}
	if cr.PH == 0.0 {
		return errEmptyPH
	}
	if cr.Timestamp == 0 {
		return errEmptyTimestamp
	}
	if cr.Moisture == 0.0 {
		return errEmptyMoisture
	}
	if cr.Temperature == 0.0 {
		return errEmptyTemperature
	}
	return
}
