package dangers

type CreateDangerRequest struct {
	Category    string
	Name        string
	Description string
	Grade       int
	UserRole    string
}
type DeleteDangerRequest struct {
	Id       int
	UserRole string
}
