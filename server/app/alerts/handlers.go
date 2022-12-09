package alerts

import "github.com/gorilla/mux"

type Service interface {
}

func RegisterRoutes(router *mux.Router, svc Service) {

}
