package infra

import (
    "encoding/json"
    "math/rand"
    "os"
    "fifaclubscup/domain"
)

type TeamJson struct {
    Name    string `json:"name"`
    Country string `json:"country"`
    Confed  string `json:"confed"`
}

type TeamLoader struct {
    SourcePath string
}

func (tl *TeamLoader) LoadTeams() ([]domain.Team, error) {
    file, err := os.Open(tl.SourcePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var teamsJson []TeamJson
    if err := json.NewDecoder(file).Decode(&teamsJson); err != nil {
        return nil, err
    }

    var teams []domain.Team
    for _, t := range teamsJson {
        teams = append(teams, domain.NewTeam(
            t.Name,
            t.Country,
            t.Confed,
            50+rand.Intn(50), // força aleatória entre 50 e 99
        ))
    }
    return teams, nil
}
