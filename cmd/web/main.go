package main

import (
	"fmt"
	"log"
	"os"
	"patient-appointment-demo-go/internal/app"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort,err := strconv.ParseInt(os.Getenv("APP_PORT"), 10, 64)
    if err != nil {
        fmt.Println("failed to parse env var APP_PORT, defaulting to 8000")
        appPort = 8000
    }

	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
    dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	app := app.New(app.ConfigWithPort(int(appPort)))

	err = app.ConnectDB(dbURL)
	if err != nil {
		log.Panicf("unable to connect to Database: %v\n", err)
	}
	defer app.CloseDB()

	fmt.Println("Succesfully connected to database")

	fmt.Printf("Starting server on port %d\n", appPort)
	err = app.Start()

	if err != nil {
		log.Fatalf("Server Error: %v", err)
	}

}

