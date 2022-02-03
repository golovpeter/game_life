package main

import (
	"github.com/golovpetr/internal/handlers/nextstep"
	"github.com/golovpetr/internal/handlers/startgame"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/startGame", startgame.StartGamePost)
	http.HandleFunc("/nextStep", nextstep.NextStep)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
