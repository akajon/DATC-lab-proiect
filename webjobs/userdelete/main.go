package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	// database connection
	// build connection string
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		"proiectdatc.database.windows.net", "CloudSA35efb96b", "22.dejlol", port, "city_danger_alert")

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
