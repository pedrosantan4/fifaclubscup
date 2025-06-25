package infra

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type HistoricTournament struct {
	Year        int    `json:"year"`
	Name        string `json:"name"`
	Champion    string `json:"champion"`
	RunnerUp    string `json:"runner_up"`
	ThirdPlace  string `json:"third_place"`
	FinalScore  string `json:"final_score"`
}

type HistoryRepository struct {
	FilePath string
}

// NewHistoryRepository initializes the repository with file path.
func NewHistoryRepository(path string) *HistoryRepository {
	return &HistoryRepository{FilePath: path}
}

// Save adds a new tournament to the history file.
func (r *HistoryRepository) Save(tournament HistoricTournament) error {
	var history []HistoricTournament

	// Cria arquivo se não existir
	if _, err := os.Stat(r.FilePath); err == nil {
		data, err := os.ReadFile(r.FilePath)
		if err == nil {
			json.Unmarshal(data, &history)
		}
	}

	history = append(history, tournament)

	// Salvar novamente
	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	// Garante diretório
	if err := os.MkdirAll(filepath.Dir(r.FilePath), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(r.FilePath, data, 0644)
}

// ListHistory prints past tournaments with top 3.
func (r *HistoryRepository) ListHistory() error {
	data, err := os.ReadFile(r.FilePath)
	if err != nil {
		return fmt.Errorf("erro ao ler histórico: %w", err)
	}

	var history []HistoricTournament
	if err := json.Unmarshal(data, &history); err != nil {
		return err
	}

	fmt.Println("\n📚 Histórico de torneios anteriores:")
	for _, t := range history {
		fmt.Printf("- %d | %s\n", t.Year, t.Name)
		fmt.Printf("   🥇 %s\n   🥈 %s\n   🥉 %s\n   🔚 Final: %s\n\n", t.Champion, t.RunnerUp, t.ThirdPlace, t.FinalScore)
	}
	return nil
}

func UnmarshalHistory(data []byte, history *[]HistoricTournament) error {
	return json.Unmarshal(data, history)
}
