package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
	"log"
	"net/http"
	"os"
	"server/app/alerts"
	"server/app/dangers"
	"server/app/users"
	"strconv"
)

func main() {
	//config
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// database connection
	// build connection string
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		os.Getenv("SERVER"), os.Getenv("USER"), os.Getenv("PASSWORD"), port, os.Getenv("DATABASE"))

	// create connection pool
	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	
	// verify connection
	ctx := context.Background()
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
	defer conn.Close()

	// repositories
	usersRepo := users.NewRepository(conn)
	dangersRepo := dangers.NewRepository(conn)
	alertsRepo := alerts.NewRepository(conn)

	// services
	userService := users.NewService(usersRepo)
	dangersService := dangers.NewService(dangersRepo)
	alertsService := alerts.NewService(alertsRepo)

	// transport
	router := mux.NewRouter()

	users.RegisterRoutes(router, userService)
	dangers.RegisterRoutes(router, dangersService)
	alerts.RegisterRoutes(router, alertsService)

	// run server
	err = http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Printf("error listening on port (port already in use?) : %#v", err)
		return
	}
}
