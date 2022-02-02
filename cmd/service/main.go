package main

import (
	"github.com/golovpetr/internal/handlers/nextstep"
	"github.com/golovpetr/internal/handlers/startgame"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/startGame", startgame.StartGamePost)
	http.HandleFunc("/nextStep", nextstep.NextStep)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
