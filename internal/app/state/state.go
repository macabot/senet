package state

import "github.com/macabot/hypp"

type Page int

const (
	GamePage Page = iota
)

type State struct {
	hypp.EmptyState
	Game    *Game
	Rotated bool
	Page    Page
}

func (s *State) Clone() *State {
	if s == nil {
		return nil
	}
	return &State{
		Game:    s.Game.Clone(),
		Rotated: s.Rotated,
	}
}
