package sweat

import "errors"

var (
	errEmptyUserID      = errors.New("User ID must be present")
	errEmptyVolume      = errors.New("Volume must be present")
	errEmptyPH          = errors.New("pH must be present")
	errEmptyTimestamp   = errors.New("Timestamp must be present")
	errEmptyMoisture    = errors.New("Moisture must be present")
	errEmptyTemperature = errors.New("Temperature must be present")
	errNoSweatId        = errors.New("Sweat is not present")
)
