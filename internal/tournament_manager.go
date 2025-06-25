// internal/tournament_manager.go
package internal

import (
    "fmt"
    "time"
    "log"

    "fifaclubscup/domain"
    "fifaclubscup/infra"
)

// TournamentManager coordinates the tournament lifecycle.
type TournamentManager struct {
    Teams []domain.Team
}

// NewTournamentManager creates a new manager with the given teams.
func NewTournamentManager(teams []domain.Team) *TournamentManager {
    return &TournamentManager{Teams: teams}
}

// RunTournament runs group stage, knockout rounds and returns the completed tournament.
func (tm *TournamentManager) RunTournament() *domain.Tournament {
    year := time.Now().Year()
    tournament := domain.NewTournament(year, "FIFA Club World Cup", tm.Teams)

    fmt.Println("🔄 Criando grupos...")
    tournament.CreateGroups()
	fmt.Println("✅ Grupos criados com sucesso!")

    fmt.Println("⚽ Simulando fase de grupos...")
    tournament.SimulateGroupStage()
	fmt.Println("✅ Fase de grupos concluída!")

    return tournament
}

// ExecuteTournamentFlow loads teams, runs the tournament and presents the results.
func ExecuteTournamentFlow() {
    // Carrega times a partir do JSON
    loader := infra.TeamLoader{SourcePath: "infra/teams.json"}
    teams, err := loader.LoadTeams()
    if err != nil {
        log.Fatalf("Erro ao carregar times: %v", err)
    }

    // Executa o torneio
    tm := NewTournamentManager(teams)
    tournament := tm.RunTournament()

    // Apresenta resultados (console + HTML)
    presenter := infra.OutputPresenter{HTMLPath: "results.html"}
    if err := presenter.Present(*tournament); err != nil {
        log.Fatalf("Erro ao apresentar resultados: %v", err)
    }
}
