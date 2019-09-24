package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/kyeett/highscore-server/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	s := service.New(db)
	if err := http.ListenAndServe(":"+port, s.Router); err != nil {
		log.Fatal(err)
	}
}
