package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kyeett/highscore-server/internal/service"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	s := service.New()
	http.ListenAndServe(":"+port, s)
}
