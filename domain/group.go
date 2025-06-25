package domain

import (
	"fmt"
	"sort"
	"sync"
	"strings"
	"time"
)

// Group represents a group with 4 teams.
type Group struct {
	Name     string
	Teams    []Team
	Matches  []*Match
	Standings map[string]*GroupStanding
}

// GroupStanding stores stats of a team within a group.
type GroupStanding struct {
	Team      Team
	Points    int
	GoalsFor  int
	GoalsAgainst int
	GoalDiff  int
}

// NewGroup creates a new group with its teams.
func NewGroup(name string, teams []Team) *Group {
	standings := make(map[string]*GroupStanding)
	for _, t := range teams {
		standings[t.ID] = &GroupStanding{Team: t}
	}

	return &Group{
		Name:     name,
		Teams:    teams,
		Matches:  []*Match{},
		Standings: standings,
	}
}

// GenerateMatches creates the 6 matches (round-robin).
func (g *Group) GenerateMatches() {
	for i := 0; i < len(g.Teams); i++ {
		for j := i + 1; j < len(g.Teams); j++ {
			match := NewMatch(g.Teams[i], g.Teams[j], GroupStage)
			g.Matches = append(g.Matches, match)
		}
	}
}

// PlayMatches runs all group matches in parallel.
func (g *Group) PlayMatches() {
	var wg sync.WaitGroup
	resultsCh := make(chan *Match, len(g.Matches))

	for _, match := range g.Matches {
		wg.Add(1)
		go func(m *Match) {
			defer wg.Done()
			m.Simulate()
			resultsCh <- m
		}(match)
	}

	wg.Wait()
	close(resultsCh)

	for result := range resultsCh {
		g.updateStandings(result)
	}
}

// updateStandings updates group table after a match result.
func (g *Group) updateStandings(m *Match) {
	home := g.Standings[m.TeamA.ID]
	away := g.Standings[m.TeamB.ID]

	home.GoalsFor += m.GoalsA
	home.GoalsAgainst += m.GoalsB
	away.GoalsFor += m.GoalsB
	away.GoalsAgainst += m.GoalsA

	home.GoalDiff = home.GoalsFor - home.GoalsAgainst
	away.GoalDiff = away.GoalsFor - away.GoalsAgainst

	if m.IsDraw {
		home.Points += 1
		away.Points += 1
	} else if m.Winner.ID == m.TeamA.ID {
		home.Points += 3
	} else {
		away.Points += 3
	}
}

// GetTopTwo returns the top 2 classified teams.
func (g *Group) GetTopTwo() []Team {
	var standingsList []*GroupStanding
	for _, s := range g.Standings {
		standingsList = append(standingsList, s)
	}

	// Sort by: points > goal diff > goals for
	sort.Slice(standingsList, func(i, j int) bool {
		a, b := standingsList[i], standingsList[j]
		if a.Points != b.Points {
			return a.Points > b.Points
		}
		if a.GoalDiff != b.GoalDiff {
			return a.GoalDiff > b.GoalDiff
		}
		return a.GoalsFor > b.GoalsFor
	})

	return []Team{standingsList[0].Team, standingsList[1].Team}
}

// PrintStandings displays group classification.
func (g *Group) PrintStandings() {
    fmt.Printf("\n\033[1;34m📊 Classificação do Grupo %s\033[0m\n", g.Name)
    fmt.Println(strings.Repeat("─", 55))
    fmt.Printf("Pos  %-22s Pts  SG  GM  GS\n", "Clube")
    fmt.Println(strings.Repeat("─", 55))

    var standingsList []*GroupStanding
    for _, s := range g.Standings {
        standingsList = append(standingsList, s)
    }

    // Sort by: points > goal diff > goals for
    sort.Slice(standingsList, func(i, j int) bool {
        a, b := standingsList[i], standingsList[j]
        if a.Points != b.Points {
            return a.Points > b.Points
        }
        if a.GoalDiff != b.GoalDiff {
            return a.GoalDiff > b.GoalDiff
        }
        return a.GoalsFor > b.GoalsFor
    })

    for i, s := range standingsList {
        fmt.Printf("%2dº  %-22s %3d %4d %4d %4d\n",
            i+1, s.Team.Name, s.Points, s.GoalDiff, s.GoalsFor, s.GoalsAgainst)
        time.Sleep(350 * time.Millisecond)
    }
    fmt.Println(strings.Repeat("─", 55))
    time.Sleep(900 * time.Millisecond)
}
