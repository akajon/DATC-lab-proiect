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
	DeleteDanger(ctx context.Context, dangerId int) error
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

// DELETEDanger delete a danger
func DELETEDanger(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var danger DeleteDangerRequest

		err := json.NewDecoder(r.Body).Decode(&danger)
		if err != nil {
			log.Print("DELETEDanger failed to decode")
			return
		}

		err = svc.DeleteDanger(r.Context(), danger.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(err.Error())
			if err != nil {
				log.Print("POSTCreateDanger failed to encode")
				return
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

func RegisterRoutes(router *mux.Router, svc Service) {
	router.Handle("/danger/create", POSTCreateDanger(svc)).Methods(http.MethodPost)
	router.Handle("/danger/delete", DELETEDanger(svc)).Methods(http.MethodDelete)
}
