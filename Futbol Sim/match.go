package main

import "math/rand"

type Match struct {
	ID        int
	HomeTeam  *Team
	AwayTeam  *Team
	HomeGoals int
	AwayGoals int
	Week      int
}

func NewMatch(home, away *Team, week int) *Match {
	return &Match{
		HomeTeam:  home,
		AwayTeam:  away,
		HomeGoals: rand.Intn(7),
		AwayGoals: rand.Intn(7),
		Week:      week,
	}
}

func (m *Match) Play() {
	if m.HomeGoals > m.AwayGoals {
		m.HomeTeam.RecordWin(m.HomeGoals, m.AwayGoals)
		m.AwayTeam.RecordLoss(m.AwayGoals, m.HomeGoals)
	} else if m.HomeGoals < m.AwayGoals {
		m.AwayTeam.RecordWin(m.AwayGoals, m.HomeGoals)
		m.HomeTeam.RecordLoss(m.HomeGoals, m.AwayGoals)
	} else {
		m.HomeTeam.RecordDraw(m.HomeGoals, m.AwayGoals)
		m.AwayTeam.RecordDraw(m.AwayGoals, m.HomeGoals)
	}
}
