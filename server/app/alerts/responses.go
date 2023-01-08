package alerts

import "time"

type AlertReward struct {
	OwnerId  int
	DangerId int
	Users    string
}

type AlertGetResponse struct {
	Id        int
	OwnerId   int
	DangerId  int
	Users     string
	Latitude  float32
	Longitude float32
	Date      *time.Time
}
