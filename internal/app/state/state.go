package state

import "github.com/macabot/hypp"

type Page int

const (
	StartPage Page = iota
	SignalingPage
	GamePage
)

type State struct {
	hypp.EmptyState
	Game *Game
	Page Page
}

func (s *State) Clone() *State {
	if s == nil {
		return nil
	}
	return &State{
		Game: s.Game.Clone(),
		Page: s.Page,
	}
}
