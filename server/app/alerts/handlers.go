package alerts

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type Service interface {
	VerifyAlert(ctx context.Context, dangerId int, latitude, longitude float32) (int, error)
	CreateAlert(ctx context.Context, userId, dangerId int, latitude, longitude float32) error
	AddUserToAlert(ctx context.Context, userId, alertId int) error
	DeleteAlert(ctx context.Context, alertId int) error
	GetAlerts(ctx context.Context) ([]AlertGetResponse, error)
}

type Claims struct {
	UserId   int
	Username string
	Role     string
	jwt.RegisteredClaims
}

func POSTAddAlert(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* c, err := r.Cookie("token")
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
		} */

		var alert CreateAlertRequest
		err := json.NewDecoder(r.Body).Decode(&alert)
		if err != nil {
			return
		}

		alertId, err := svc.VerifyAlert(r.Context(), alert.DangerId, alert.Latitude, alert.Longitude)
		if err != sql.ErrNoRows && err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if alertId == 0 {
			err := svc.CreateAlert(r.Context(), alert.UserId, alert.DangerId, alert.Latitude, alert.Longitude)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			return
		}

		err = svc.AddUserToAlert(r.Context(), alert.UserId, alertId)
		if err != nil {
			fmt.Println(err)
			if err.Error() == "user already reported this alert" {
				w.WriteHeader(http.StatusConflict)
				json.NewEncoder(w).Encode(&struct{ Error string }{Error: err.Error()})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func DELETEAlert(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* c, err := r.Cookie("token")
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
		} */

		var alert DeleteAlertRequest

		err := json.NewDecoder(r.Body).Decode(&alert)
		if err != nil {
			return
		}

		if alert.UserRole != "ADMIN" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = svc.DeleteAlert(r.Context(), alert.Id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func GETAlerts(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* c, err := r.Cookie("token")
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
		} */

		alerts, err := svc.GetAlerts(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(&alerts)
	})
}

func RegisterRoutes(router *mux.Router, svc Service) {
	router.Handle("/alert/add", POSTAddAlert(svc)).Methods(http.MethodPost)
	router.Handle("/alert/delete", DELETEAlert(svc)).Methods(http.MethodDelete)
	router.Handle("/alert", GETAlerts(svc)).Methods(http.MethodGet)
}
