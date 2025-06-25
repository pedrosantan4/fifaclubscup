package domain

import (
	"fmt"
	"math/rand"
	"time"
)

// MatchStage defines the stage of the tournament.
type MatchStage string

const (
	GroupStage     MatchStage = "Group"
	RoundOf16      MatchStage = "Round of 16"
	QuarterFinal   MatchStage = "Quarter-final"
	SemiFinal      MatchStage = "Semi-final"
	Final          MatchStage = "Final"
	ThirdPlacePlay MatchStage = "Third Place"
)

// Match holds information about a game between two teams.
type Match struct {
	TeamA     Team
	TeamB     Team
	GoalsA    int
	GoalsB    int
	Stage     MatchStage
	IsDraw    bool
	IsPlayed  bool
	Winner    *Team
	Loser     *Team
	MatchID   string
}

// NewMatch creates a new Match instance.
func NewMatch(teamA, teamB Team, stage MatchStage) *Match {
	return &Match{
		TeamA:   teamA,
		TeamB:   teamB,
		Stage:   stage,
		MatchID: generateMatchID(teamA, teamB),
	}
}

// Simulate runs the match logic based on team strength.
func (m *Match) Simulate() {
	rand.Seed(time.Now().UnixNano())

	// Gol baseado na força dos times
	baseGols := func(force int) int {
		return rand.Intn(3) + force/35
	}

	g1 := baseGols(m.TeamA.Force)
	g2 := baseGols(m.TeamB.Force)

	// desempate forçado no mata-mata
	if g1 == g2 && m.Stage != GroupStage {
		// força pênaltis ou prorrogação (simulação simples)
		if rand.Intn(2) == 0 {
			g1++
		} else {
			g2++
		}
	}

	m.GoalsA = g1
	m.GoalsB = g2

	if g1 == g2 {
		m.IsDraw = true
		m.Winner = nil
		m.Loser = nil
	} else {
		m.IsDraw = false
		if g1 > g2 {
			m.Winner = &m.TeamA
			m.Loser = &m.TeamB
		} else {
			m.Winner = &m.TeamB
			m.Loser = &m.TeamA
		}
	}
}


// Summary returns a formatted string with the result.
func (m *Match) Summary() string {
	if !m.IsPlayed {
		return fmt.Sprintf("%s vs %s (not played)", m.TeamA.Name, m.TeamB.Name)
	}
	return fmt.Sprintf("[%s] %s %d x %d %s", m.Stage, m.TeamA.Name, m.GoalsA, m.GoalsB, m.TeamB.Name)
}

// generateMatchID creates a unique string identifier for a match.
func generateMatchID(teamA, teamB Team) string {
	return fmt.Sprintf("%s_vs_%s_%d", teamA.Name, teamB.Name, time.Now().UnixNano())
}
