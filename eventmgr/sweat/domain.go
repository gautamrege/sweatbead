package sweat

type createRequest struct {
	UserId      int64   `json: user_id`
	Volume      float32 `json:volume`
	PH          float32 `json:pH`
	Timestamp   int64   `json:timestamp`
	Moisture    float32 `json:moisture`
	Temperature float32 `json:temperature`
}

func (cr createRequest) Validate() (err error) {
	if cr.UserId == 0 {
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
