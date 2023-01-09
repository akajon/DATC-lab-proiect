package alerts

type CreateAlertRequest struct {
	DangerId  int
	Latitude  float32
	Longitude float32
	UserId    int
}

type DeleteAlertRequest struct {
	Id       int
	UserRole string
}
