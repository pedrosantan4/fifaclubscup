package main

import (
	"fifaclubscup/domain"
	"fifaclubscup/infra"
	"fifaclubscup/internal"
	"fmt"
	"os"
	"strings"
	"time"
)

func printSeparator() {
    fmt.Println("\n" + strings.Repeat("═", 50) + "\n")
}

func showLastChampion(historyPath string) {
    file, err := os.ReadFile(historyPath)
    if err != nil || len(file) == 0 {
        fmt.Println("Nenhum campeão anterior registrado.")
        return
    }
    var history []infra.HistoricTournament
    if err := infra.UnmarshalHistory(file, &history); err != nil || len(history) == 0 {
        fmt.Println("Nenhum campeão anterior registrado.")
        return
    }
    last := history[len(history)-1]
    fmt.Printf("🏆 Último campeão: \033[1;36m%s\033[0m (%d)\n", last.Champion, last.Year)
}

func showGroups(groups []*domain.Group) {
    for _, group := range groups {
        fmt.Printf("\n\033[1;33mGrupo %s\033[0m:\n", group.Name)
        time.Sleep(350 * time.Millisecond)
        for _, team := range group.Teams {
            fmt.Printf("   • \033[1;37m%s\033[0m (\033[0;36m%s\033[0m)\n", team.Name, team.Country)
            time.Sleep(100 * time.Millisecond)
        }
        time.Sleep(200 * time.Millisecond)
    }
}

func showGroupStandings(groups []*domain.Group) {
    for _, group := range groups {
        group.PrintStandings()
        time.Sleep(900 * time.Millisecond)
    }
}

func showKnockoutMatches(finals []*domain.Match) {
    stages := []string{"Round of 16", "Quarter-final", "Semi-final", "Final", "Third Place"}
    stageNames := map[string]string{
        "Round of 16":   "Oitavas de Final",
        "Quarter-final": "Quartas de Final",
        "Semi-final":    "Semifinais",
        "Final":         "Grande Final",
        "Third Place":   "Disputa 3º Lugar",
    }
    for _, stage := range stages {
        var matches []*domain.Match
        for _, m := range finals {
            if string(m.Stage) == stage {
                matches = append(matches, m)
            }
        }
        if len(matches) > 0 {
            fmt.Printf("\n\033[1;35m%s\033[0m\n", stageNames[stage])
            time.Sleep(700 * time.Millisecond)
            for _, m := range matches {
                fmt.Printf("   %s x %s", m.TeamA.Name, m.TeamB.Name)
                time.Sleep(600 * time.Millisecond)
                if m.IsPlayed {
                    fmt.Printf("  →  \033[1;32m%d\033[0m x \033[1;31m%d\033[0m", m.GoalsA, m.GoalsB)
                    if m.Winner != nil {
                        fmt.Printf("  |  🏅 Vencedor: \033[1;36m%s\033[0m", m.Winner.Name)
                    }
                }
                fmt.Println()
                time.Sleep(700 * time.Millisecond)
            }
            time.Sleep(900 * time.Millisecond)
        }
    }
}

func main() {
    fmt.Println("\033[1;34m🌍 Bem-vindo ao Mundial de Clubes FIFA!\033[0m")
    time.Sleep(1200 * time.Millisecond)

    historyPath := "infra/data/history.json"
    printSeparator()
    fmt.Println("🔎 Consultando histórico recente...")
    time.Sleep(900 * time.Millisecond)
    showLastChampion(historyPath)
    time.Sleep(1800 * time.Millisecond)

    printSeparator()
    fmt.Println("🔄 Carregando clubes participantes...")
    time.Sleep(1200 * time.Millisecond)
    loader := infra.TeamLoader{SourcePath: "infra/teams.json"}
    teams, err := loader.LoadTeams()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Foram encontrados \033[1;32m%d\033[0m clubes!\n", len(teams))
    time.Sleep(1200 * time.Millisecond)

    manager := internal.NewTournamentManager(teams)
    tournament := manager.RunTournament()

    printSeparator()
    fmt.Println("🏁 \033[1;33mApresentação dos Grupos\033[0m")
    time.Sleep(1200 * time.Millisecond)
    showGroups(tournament.Groups)
    time.Sleep(1200 * time.Millisecond)

    printSeparator()
    fmt.Println("⚽ \033[1;36mSimulando jogos da fase de grupos...\033[0m")
    time.Sleep(1800 * time.Millisecond)
    tournament.SimulateGroupStage()
    fmt.Println("✅ Fase de grupos concluída!")
    time.Sleep(1200 * time.Millisecond)

    printSeparator()
    fmt.Println("📊 \033[1;36mClassificação dos Grupos\033[0m")
    time.Sleep(1000 * time.Millisecond)
    showGroupStandings(tournament.Groups)
    time.Sleep(1200 * time.Millisecond)

    printSeparator()
    fmt.Println("🏆 \033[1;35mFase Eliminatória\033[0m")
    time.Sleep(1200 * time.Millisecond)
    tournament.SimulateKnockout()
    tournament.SimulateThirdPlaceGame()
    fmt.Println("✅ Fase eliminatória concluída!")
    time.Sleep(1200 * time.Millisecond)

    printSeparator()
    fmt.Println("🔢 \033[1;33mResultados dos confrontos eliminatórios:\033[0m")
    showKnockoutMatches(tournament.Finals)
    time.Sleep(1200 * time.Millisecond)

    // Salva histórico apenas do torneio atual
    repo := infra.NewHistoryRepository(historyPath)
    finalMatch := tournament.Finals[len(tournament.Finals)-1]
    historic := infra.HistoricTournament{
        Year:       tournament.Year,
        Name:       tournament.Name,
        Champion:   tournament.Champion.Name,
        RunnerUp:   tournament.RunnerUp.Name,
        ThirdPlace: tournament.ThirdPlace.Name,
        FinalScore: fmt.Sprintf("%d x %d", finalMatch.GoalsA, finalMatch.GoalsB),
    }
    _ = repo.Save(historic)

    printSeparator()
    fmt.Println("🎬 \033[1;34mResumo do Torneio\033[0m")
    time.Sleep(1000 * time.Millisecond)
    tournament.Summary()
    time.Sleep(1200 * time.Millisecond)

    fmt.Printf("\n🏅 \033[1;32mCampeão atual: %s\033[0m\n", tournament.Champion.Name)
    fmt.Println("🏁 \033[1;34mFim do torneio! Obrigado por acompanhar.\033[0m")

    // Atualiza o HTML ao final
    presenter := &infra.OutputPresenter{HTMLPath: "output.html"}
    presenter.Present(*tournament)
}