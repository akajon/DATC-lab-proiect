package dangers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type Service interface {
	CreateDanger(ctx context.Context, category, name, description string, grade int) error
	DeleteDanger(ctx context.Context, dangerId int) error
}

type Claims struct {
	UserId   int
	Username string
	jwt.RegisteredClaims
}

// POSTCreateDanger create a danger
func POSTCreateDanger(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tknStr := c.Value
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var danger CreateDangerRequest

		err = json.NewDecoder(r.Body).Decode(&danger)
		if err != nil {
			log.Print("POSTCreateDanger failed to decode")
			return
		}

		err = svc.CreateDanger(r.Context(), danger.Category, danger.Name, danger.Description, danger.Grade)
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
	})
}

// DELETEDanger delete a danger
func DELETEDanger(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tknStr := c.Value
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var danger DeleteDangerRequest

		err = json.NewDecoder(r.Body).Decode(&danger)
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
	})
}

func RegisterRoutes(router *mux.Router, svc Service) {
	router.Handle("/danger/create", POSTCreateDanger(svc)).Methods(http.MethodPost)
	router.Handle("/danger/delete", DELETEDanger(svc)).Methods(http.MethodDelete)
}
