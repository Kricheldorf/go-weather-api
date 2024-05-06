package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-weather-api/internal/api/handlers"
	"go-weather-api/internal/api/models"
	"log"
	"net/http"
	"os"
)

func listenAndServe(router *mux.Router) {
	port := fmt.Sprintf("%s:%s", os.Getenv("BASE_APP_HOST"), os.Getenv("DEFAULT_PORT"))

	err := http.ListenAndServe(port, router)
	if err != nil {
		fmt.Printf("Port %s is already in use. Trying another port...\n", port)
		port := fmt.Sprintf("%s:%s", os.Getenv("BASE_APP_HOST"), os.Getenv("RETRY_PORT"))
		err = http.ListenAndServe(port, router)
		if err != nil {
			fmt.Printf("Failed to start server on port %s: %v\n", port, err)
			return
		}
	}

	fmt.Printf("Server started on port %s\n", port)
}

func initDB() {
	var err error

	postgresURI := os.Getenv("DATABASE_URL")
	models.DB, err = sql.Open("postgres", postgresURI)
	if err != nil {
		log.Fatal(err)
	}
	err = models.DB.Ping()
	if err != nil {
		err = models.DB.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}

	fmt.Println("Connected to database")
	return
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	initDB()

	router := handlers.SetupRoutes()

	listenAndServe(router)
}
