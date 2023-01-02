package dangers

type CreateDangerRequest struct {
	Category    string
	Name        string
	Description string
	Grade       int
}
type DeleteDangerRequest struct {
	Id int
}
