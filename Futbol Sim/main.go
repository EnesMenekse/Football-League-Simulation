package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

var league *League

func main() {
	initDB()
	rand.Seed(time.Now().UnixNano())

	teamNames := []string{"Liverpool", "Manchester United", "Manchester City", "Chelsea"}
	league = NewLeague(teamNames)

	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	http.HandleFunc("/simulate", simulateHandler)
	http.HandleFunc("/standings", standingsHandler)
	http.HandleFunc("/next-week", nextWeekHandler)
	http.HandleFunc("/matches", matchesHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
