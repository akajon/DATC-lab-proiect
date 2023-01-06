package alerts

type CreateAlertRequest struct {
	DangerId  int
	Latitude  float32
	Longitude float32
}

type DeleteAlertRequest struct {
	Id int
}
