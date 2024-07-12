package main

type Team struct {
	ID           int
	Name         string
	Points       int
	Wins         int
	Draws        int
	Losses       int
	GoalsFor     int
	GoalsAgainst int
}

func NewTeam(name string) *Team {
	return &Team{Name: name}
}

func (t *Team) RecordWin(goalsFor, goalsAgainst int) {
	t.Wins++
	t.Points += 3
	t.GoalsFor += goalsFor
	t.GoalsAgainst += goalsAgainst
	saveTeamStats(t)
}

func (t *Team) RecordDraw(goalsFor, goalsAgainst int) {
	t.Draws++
	t.Points++
	t.GoalsFor += goalsFor
	t.GoalsAgainst += goalsAgainst
	saveTeamStats(t)
}

func (t *Team) RecordLoss(goalsFor, goalsAgainst int) {
	t.Losses++
	t.GoalsFor += goalsFor
	t.GoalsAgainst += goalsAgainst
	saveTeamStats(t)
}
