package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "user=postgres password=1234 dbname=FootballLeague sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	clearTables()
}

func clearTables() {
	_, err := db.Exec("DELETE FROM matches")
	if err != nil {
		log.Println("Error clearing matches table:", err)
	}

	_, err = db.Exec("DELETE FROM teams")
	if err != nil {
		log.Println("Error clearing teams table:", err)
	}

	_, err = db.Exec("ALTER SEQUENCE teams_id_seq RESTART WITH 1")
	if err != nil {
		log.Println("Error resetting team ID sequence:", err)
	}

	_, err = db.Exec("ALTER SEQUENCE matches_id_seq RESTART WITH 1")
	if err != nil {
		log.Println("Error resetting match ID sequence:", err)
	}
}

func saveTeamStats(team *Team) {
	_, err := db.Exec(`
        UPDATE teams
        SET points = $1, wins = $2, draws = $3, losses = $4, goals_for = $5, goals_against = $6
        WHERE name = $7`,
		team.Points, team.Wins, team.Draws, team.Losses, team.GoalsFor, team.GoalsAgainst, team.Name)
	if err != nil {
		log.Println("Error updating team stats:", err)
	}
}

func saveMatch(match *Match) {
	_, err := db.Exec(`
        INSERT INTO matches (home_team_id, away_team_id, home_goals, away_goals, week)
        VALUES ((SELECT id FROM teams WHERE name = $1), (SELECT id FROM teams WHERE name = $2), $3, $4, $5)`,
		match.HomeTeam.Name, match.AwayTeam.Name, match.HomeGoals, match.AwayGoals, match.Week)
	if err != nil {
		log.Println("Error saving match:", err)
	}
}

func insertTeam(team *Team) {
	if !teamExists(team.Name) {
		_, err := db.Exec(`
			INSERT INTO teams (name, points, wins, draws, losses, goals_for, goals_against)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			team.Name, team.Points, team.Wins, team.Draws, team.Losses, team.GoalsFor, team.GoalsAgainst)
		if err != nil {
			log.Println("Error inserting team:", err)
		}
	}
}

func teamExists(name string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM teams WHERE name=$1)`
	err := db.QueryRow(query, name).Scan(&exists)
	if err != nil {
		log.Println("Error checking if team exists:", err)
	}
	return exists
}

func saveMatches(matches []*Match) {
	for _, match := range matches {
		saveMatch(match)
	}
}
