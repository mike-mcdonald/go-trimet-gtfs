package types

// ReturnValue is the root value for the vehicle feed
type ReturnValue struct {
	ResultSet ResultSet `json:"resultSet"`
}

// ResultSet is the base return type for the vehicle feed
type ResultSet struct {
	QueryTime int64     `json:"queryTime"`
	Vehicle   []Vehicle `json:"vehicle"`
}

// Error is returned when there was a problem with the web service
type Error struct {
	Content string `json:"content"`
}

// Vehicle is the object that describes a vehicle position
type Vehicle struct {
	VehicleID             int     `json:"vehicleID"`
	BlockID               int     `json:"blockID"`
	Bearing               int16   `json:"bearing"`
	ServiceDate           int64   `json:"serviceDate"`
	LocationInScheduleDay int     `json:"locationInScheduleDay"`
	Time                  int64   `json:"time"`
	Expires               int64   `json:"expires"`
	Longitude             float32 `json:"longitude"`
	Latitude              float32 `json:"latitude"`
	RouteNumber           int16   `json:"routeNumber"`
	Direction             int8    `json:"direction"`
	TripID                string  `json:"tripID"`
	NewTrip               bool    `json:"newTrip"`
	Delay                 int16   `json:"delay"`
	MessageCode           int     `json:"messageCode"`
	SignMessage           string  `json:"signMessage"`
	SignMessageLong       string  `json:"signMessageLong"`
	NextLocID             int     `json:"nextLocID"`
	NextStopSeq           int     `json:"nextStopSeq"`
	LastLocID             int     `json:"lastLocID"`
	LastStopSeq           int     `json:"lastStopSeq"`
	Garage                string  `json:"garage"`
	ExtrablockID          int     `json:"extrablockID"`
	OffRoute              bool    `json:"offRoute"`
	InCongestion          bool    `json:"inCongestion"`
	LoadPercentage        int     `json:"loadPercentage"`
	Source                string  `json:"source"`
}
