---
title: Football League Simulation Documentation
---

# Project Name

Football League Simulation

# Project Description

This project simulates a football league where teams like Liverpool,
Manchester United, Manchester City, and Chelsea play random matches. The
teams' scores are updated based on match results, and league standings
are determined. The project is developed using the Go programming
language and PostgreSQL database.

# Installation Instructions

1\. PostgreSQL Setup and Database Creation

CREATE DATABASE FootballLeague;  
  
CREATE TABLE teams (  
id SERIAL PRIMARY KEY,  
name VARCHAR(50) NOT NULL,  
points INT DEFAULT 0,  
wins INT DEFAULT 0,  
draws INT DEFAULT 0,  
losses INT DEFAULT 0,  
goals_for INT DEFAULT 0,  
goals_against INT DEFAULT 0  
);  
CREATE TABLE matches (  
id SERIAL PRIMARY KEY,  
home_team_id INT REFERENCES teams(id),  
away_team_id INT REFERENCES teams(id),  
home_goals INT,  
away_goals INT,  
week INT  
);

2\. Download Project Files

git clone https://github.com/EnesMenekse/Football-League-Simulation.git

cd Football-League-Simulation

3\. Install Required Modules

The project directory contains go.mod and go.sum files which manage the
project dependencies. To install the necessary dependencies, run:

go mod tidy

4\. Start the Server

go run main.go database.go team.go match.go league.go handlers.go
model.go

# Usage Instructions

After starting the server, open your browser and go to
http://localhost:8080 or use Postman to interact with the API endpoints.

# File Descriptions

## main.go

The entry point of the application. It initializes the database
connection, sets the seed for the random number generator, creates the
league, and starts the HTTP server.

## database.go

Used to connect to the PostgreSQL database and perform basic database
operations. Connects to the database and checks the connection. Clears
and resets the tables.

## model.go

Contains database operations and SQL queries. Updates team statistics.
Saves match results to the database. Adds new teams to the database.
Checks if a team exists in the database.

## team.go

Defines the team structure and updates team statistics. Defines the Team
structure. Updates team statistics based on win, draw, or loss.

## match.go

Defines the match structure and simulates the matches. Defines the Match
structure. Determines the results of the matches and updates team
statistics.

## league.go

Defines the league structure and manages the scheduling and playing of
matches in the league. Defines the League structure. Schedules league
matches. Plays weekly matches and saves the results.

## handlers.go

Defines handler functions for the HTTP endpoints. These functions
receive HTTP requests and perform the relevant operations. Simulates the
entire season and returns the results. Simulates the next week's matches
and returns the new match results. Returns all match results. Returns
the current league standings.

# API Endpoints

POST /simulate: Simulates the entire season.  
GET /standings: Returns the current league standings.  
POST /next-week: Simulates the next week's matches.  
GET /matches: Returns all match results.

# Example Requests

Start Simulation:

POST http://localhost:8080/simulate

Get Current Standings:

GET http://localhost:8080/standings

Simulate Next Week's Matches:

POST http://localhost:8080/next-week

Get All Match Results:

GET http://localhost:8080/matches
