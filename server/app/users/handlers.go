package users

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(ctx context.Context, firstName, lastName, email, password string) error
	UpdateDeleteDate(ctx context.Context, userId int) error
	GetUserPasswordAndId(ctx context.Context, username string) (string, int, error)
	GetRole(ctx context.Context, userId int) (string, error)
}

type Claims struct {
	UserId   int
	Username string
	Role     string
	jwt.RegisteredClaims
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

		err = svc.CreateUser(r.Context(), user.FirstName, user.LastName, user.Email, user.Password)
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
	})
}

// PUTUpdateDeleteDate updates user deletion date (30 days from today) to be deleted by web job worker after 30days
func PUTUpdateDeleteDate(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = svc.UpdateDeleteDate(r.Context(), claims.UserId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(err.Error())
			if err != nil {
				log.Print("PUTUpdateDeleteDate failed to encode")
				return
			}
			return
		}
	})
}

func POSTSignIn(svc Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials

		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		hashedPassword, userId, err := svc.GetUserPasswordAndId(r.Context(), creds.Username)

		err2 := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))

		if err != nil || err2 != nil {
			fmt.Println(err)
			fmt.Println(err2)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		role, err := svc.GetRole(r.Context(), userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		expirationTime := time.Now().Add(30 * time.Minute)
		claims := &Claims{
			UserId:   userId,
			Username: creds.Username,
			Role:     role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	})
}

func RegisterRoutes(router *mux.Router, svc Service) {
	router.Handle("/user/register", POSTCreateUser(svc)).Methods(http.MethodPost)
	router.Handle("/user/delete", PUTUpdateDeleteDate(svc)).Methods(http.MethodPut)
	router.Handle("/signin", POSTSignIn(svc)).Methods(http.MethodPost)
}
