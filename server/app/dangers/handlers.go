package dangers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Service interface {
	CreateDanger(ctx context.Context, category, name, description string, grade int) (*CreateDangerResponse, error)
}

// POSTCreateDanger create a danger
func POSTCreateDanger(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var danger CreateDangerRequest

		err := json.NewDecoder(r.Body).Decode(&danger)
		if err != nil {
			log.Print("POSTCreateDanger failed to decode")
			return
		}

		createdDanger, err := svc.CreateDanger(r.Context(), danger.Category, danger.Name, danger.Description, danger.Grade)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(err.Error())
			if err != nil {
				log.Print("POSTCreateDanger failed to encode")
				return
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(createdDanger)
		if err != nil {
			log.Print("POSTCreateDanger failed to encode")
			return
		}
	})
}

func RegisterRoutes(router *mux.Router, svc Service) {
	router.Handle("/danger/create", POSTCreateDanger(svc)).Methods(http.MethodPost)
}
