package users

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Service interface {
	CreateUser(ctx context.Context, firstName, lastName, email, password string) (*CreateUserResponse, error)
	UpdateDeleteDate(ctx context.Context, userId int) (*UpdateDeleteDateResponse, error)
}

// POSTCreateUser create user account
func POSTCreateUser(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Print("POSTCreateUser failed to decode")
			return
		}

		newUser, err := svc.CreateUser(r.Context(), user.FirstName, user.LastName, user.Email, user.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(err.Error())
			if err != nil {
				log.Print("POSTCreateUser failed to encode")
				return
			}
			return
		}

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(newUser)
		if err != nil {
			log.Print("POSTCreateUser failed to encode")
			return
		}
	})
}

// PUTUpdateDeleteDate updates user deletion date (30 days from today) to be deleted by web job worker after 30days
func PUTUpdateDeleteDate(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user UpdateDeleteDateRequest

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Print("PUTUpdateDeleteDate failed to decode")
			return
		}

		updatedUser, err := svc.UpdateDeleteDate(r.Context(), user.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(err.Error())
			if err != nil {
				log.Print("PUTUpdateDeleteDate failed to encode")
				return
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(updatedUser)
		if err != nil {
			log.Print("PUTUpdateDeleteDate failed to encode")
			return
		}
	})
}

func RegisterRoutes(router *mux.Router, svc Service) {
	router.Handle("/user/create", POSTCreateUser(svc)).Methods(http.MethodPost)
	router.Handle("/user/delete", PUTUpdateDeleteDate(svc)).Methods(http.MethodPut)
}
