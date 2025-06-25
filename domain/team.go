package domain

import (
	"fmt"

	"github.com/google/uuid"
)

// Team represents a football club in the tournament.
type Team struct {
	ID            string
	Name          string
	Country       string
	Confederation string
	Strength      int // Used to simulate match outcomes
	Force		 int // Deprecated: Use Strength instead
}

// NewTeam creates a new Team with a unique ID.
func NewTeam(name, country, confederation string, strength int) Team {
	return Team{
		ID:            uuid.New().String(),
		Name:          name,
		Country:       country,
		Confederation: confederation,
		Strength:      strength,
	}
}

// String returns a human-readable string representation of the team.
func (t Team) String() string {
	return fmt.Sprintf("%s (%s - %s)", t.Name, t.Country, t.Confederation)
}
