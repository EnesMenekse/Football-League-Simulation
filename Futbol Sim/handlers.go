package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func simulateHandler(w http.ResponseWriter, r *http.Request) {
	for league.CurrentWeek < len(league.Weeks) {
		league.PlayNextWeek()
	}
	saveMatches(league.Matches)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(league)
}

func nextWeekHandler(w http.ResponseWriter, r *http.Request) {
	league.PlayNextWeek()

	log.Printf("Current week: %d", league.CurrentWeek)
	log.Println("New matches played:")

	newMatches := league.Weeks[league.CurrentWeek-1]
	for _, match := range newMatches {
		saveMatch(match)
		log.Printf("Match: %s vs %s - %d:%d", match.HomeTeam.Name, match.AwayTeam.Name, match.HomeGoals, match.AwayGoals)
	}

	rows, err := db.Query(`SELECT id, name, points, wins, draws, losses, goals_for, goals_against FROM teams`)
	if err != nil {
		http.Error(w, "Unable to query standings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var teams []*Team
	for rows.Next() {
		var team Team
		if err := rows.Scan(&team.ID, &team.Name, &team.Points, &team.Wins, &team.Draws, &team.Losses, &team.GoalsFor, &team.GoalsAgainst); err != nil {
			http.Error(w, "Unable to scan standings", http.StatusInternalServerError)
			return
		}
		teams = append(teams, &team)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error occurred during iteration", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

func matchesHandler(w http.ResponseWriter, r *http.Request) {
	type MatchScore struct {
		HomeTeam  string `json:"home_team"`
		AwayTeam  string `json:"away_team"`
		HomeGoals int    `json:"home_goals"`
		AwayGoals int    `json:"away_goals"`
	}

	rows, err := db.Query(`
		SELECT
			ht.name as home_team,
			at.name as away_team,
			m.home_goals,
			m.away_goals
		FROM matches m
		JOIN teams ht ON m.home_team_id = ht.id
		JOIN teams at ON m.away_team_id = at.id
	`)
	if err != nil {
		http.Error(w, "Unable to query matches", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var scores []MatchScore
	for rows.Next() {
		var score MatchScore
		if err := rows.Scan(&score.HomeTeam, &score.AwayTeam, &score.HomeGoals, &score.AwayGoals); err != nil {
			http.Error(w, "Unable to scan matches", http.StatusInternalServerError)
			return
		}
		scores = append(scores, score)
		log.Printf("Fetched match: HomeTeam=%s, AwayTeam=%s, HomeGoals=%d, AwayGoals=%d", score.HomeTeam, score.AwayTeam, score.HomeGoals, score.AwayGoals)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error occurred during iteration", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scores)
}

func standingsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT id, name, points, wins, draws, losses, goals_for, goals_against FROM teams`)
	if err != nil {
		http.Error(w, "Unable to query standings", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var teams []*Team
	for rows.Next() {
		var team Team
		if err := rows.Scan(&team.ID, &team.Name, &team.Points, &team.Wins, &team.Draws, &team.Losses, &team.GoalsFor, &team.GoalsAgainst); err != nil {
			http.Error(w, "Unable to scan standings", http.StatusInternalServerError)
			return
		}
		teams = append(teams, &team)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error occurred during iteration", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}
