package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
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
	fmt.Printf("Database Connected!")
	defer conn.Close()

	conn.Exec(`DELETE FROM dbo.users WHERE deletion_date = @deletion_date`, sql.Named("deletion_date", time.Now()))
}
