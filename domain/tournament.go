package domain

import (
	"math/rand"
	"time"
	"fmt"
)

type Tournament struct {
	Year        int
	Name        string
	Teams       []Team
	Groups      []*Group
	Finals      []*Match
	Champion    *Team
	RunnerUp    *Team
	ThirdPlace  *Team
}

// NewTournament creates a new tournament instance.
func NewTournament(year int, name string, teams []Team) *Tournament {
	return &Tournament{
		Year:  year,
		Name:  name,
		Teams: teams,
	}
}

// CreateGroups divides the teams into 8 groups of 4.
func (t *Tournament) CreateGroups() {
	rand.Seed(time.Now().UnixNano())

	// Embaralhar os times
	shuffled := make([]Team, len(t.Teams))
	copy(shuffled, t.Teams)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	for i := 0; i < 8; i++ {
		groupTeams := shuffled[i*4 : (i+1)*4]
		group := NewGroup(string('A'+i), groupTeams)
		group.GenerateMatches()
		t.Groups = append(t.Groups, group)
	}
}

// SimulateGroupStage simula todos os grupos
func (t *Tournament) SimulateGroupStage() {
	for _, group := range t.Groups {
		group.PlayMatches()
	}
}

// GetQualifiedTeams retorna os dois primeiros de cada grupo (total 16)
func (t *Tournament) GetQualifiedTeams() []Team {
	var qualified []Team
	for _, group := range t.Groups {
		top := group.GetTopTwo()
		qualified = append(qualified, top...)
	}
	return qualified
}

// SimulateKnockout simula as fases eliminatórias até a final
func (t *Tournament) SimulateKnockout() {
	qualified := t.GetQualifiedTeams()

	for len(qualified) > 1 {
		var stage MatchStage
		switch len(qualified) {
		case 16:
			stage = RoundOf16
		case 8:
			stage = QuarterFinal
		case 4:
			stage = SemiFinal
		case 2:
			stage = Final
		}

		var nextRound []Team
		var roundMatches []*Match

		for i := 0; i < len(qualified); i += 2 {
			match := NewMatch(qualified[i], qualified[i+1], stage)
			match.Simulate()
			roundMatches = append(roundMatches, match)
			nextRound = append(nextRound, *match.Winner)
		}

		t.Finals = append(t.Finals, roundMatches...)

		if len(nextRound) == 1 {
			t.Champion = &nextRound[0]
			lastMatch := roundMatches[len(roundMatches)-1]
			t.RunnerUp = lastMatch.Loser
		}

		qualified = nextRound
	}
}

// SimulateThirdPlaceGame cria e simula a disputa do 3º lugar
func (t *Tournament) SimulateThirdPlaceGame() {
	// pegar semifinalistas perdedores
	var semifinalLosers []*Team
	for _, m := range t.Finals {
		if m.Stage == SemiFinal {
			semifinalLosers = append(semifinalLosers, m.Loser)
		}
	}
	if len(semifinalLosers) == 2 {
		match := NewMatch(*semifinalLosers[0], *semifinalLosers[1], ThirdPlacePlay)
		match.Simulate()
		t.Finals = append(t.Finals, match)
		t.ThirdPlace = match.Winner
	}
}


func (t *Tournament) Summary() {
	fmt.Printf("\n🏆 %s - %d\n", t.Name, t.Year)
	fmt.Printf("🥇 Campeão: %s\n", t.Champion.Name)
	fmt.Printf("🥈 Vice: %s\n", t.RunnerUp.Name)
	fmt.Printf("🥉 Terceiro: %s\n", t.ThirdPlace.Name)
}
