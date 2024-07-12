package main

type League struct {
	Teams       []*Team
	Matches     []*Match
	CurrentWeek int
	Weeks       [][]*Match
}

func NewLeague(teamNames []string) *League {
	league := &League{}
	for _, name := range teamNames {
		team := NewTeam(name)
		league.Teams = append(league.Teams, team)
		insertTeam(team)
	}
	league.ScheduleMatches()
	return league
}

func (l *League) ScheduleMatches() {
	numTeams := len(l.Teams)
	totalWeeks := (numTeams - 1) * 2
	l.Weeks = make([][]*Match, totalWeeks)

	for round := 0; round < totalWeeks/2; round++ {
		for i := 0; i < numTeams/2; i++ {
			home := l.Teams[i]
			away := l.Teams[numTeams-i-1]

			homeMatch := NewMatch(home, away, round)
			l.Weeks[round] = append(l.Weeks[round], homeMatch)

			awayMatch := NewMatch(away, home, round+totalWeeks/2)
			l.Weeks[round+totalWeeks/2] = append(l.Weeks[round+totalWeeks/2], awayMatch)
		}

		lastTeam := l.Teams[numTeams-1]
		copy(l.Teams[2:], l.Teams[1:numTeams-1])
		l.Teams[1] = lastTeam
	}
}

func (l *League) PlayNextWeek() {
	if l.CurrentWeek >= len(l.Weeks) {
		return
	}
	weekMatches := l.Weeks[l.CurrentWeek]
	for _, match := range weekMatches {
		match.Play()
	}
	l.Matches = append(l.Matches, weekMatches...)
	l.CurrentWeek++
}
