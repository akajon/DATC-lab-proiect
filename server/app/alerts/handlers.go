package alerts

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type Service interface {
	VerifyAlert(ctx context.Context, dangerId int, latitude, longitude float32) (error, int)
	CreateAlert(ctx context.Context, userId, dangerId int, latitude, longitude float32) error
	AddUserToAlert(ctx context.Context, userId, alertId int) error
}

type Claims struct {
	UserId   int
	Username string
	jwt.RegisteredClaims
}

func POSTAddAlert(svc Service) http.Handler {
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

		var alert CreateAlertRequest
		err = json.NewDecoder(r.Body).Decode(&alert)
		if err != nil {
			return
		}

		err, alertId := svc.VerifyAlert(r.Context(), alert.DangerId, alert.Latitude, alert.Longitude)
		if err != sql.ErrNoRows && err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if alertId == 0 {
			err := svc.CreateAlert(r.Context(), claims.UserId, alert.DangerId, alert.Latitude, alert.Longitude)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			return
		}

		err = svc.AddUserToAlert(r.Context(), claims.UserId, alertId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func RegisterRoutes(router *mux.Router, svc Service) {
	router.Handle("/alert/add", POSTAddAlert(svc)).Methods(http.MethodPost)
}
