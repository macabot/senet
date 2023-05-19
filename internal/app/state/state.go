package state

import "github.com/macabot/hypp"

type State struct {
	hypp.EmptyState
	Game *Game
}

func (s *State) Clone() *State {
	if s == nil {
		return nil
	}
	return &State{
		Game: s.Game.Clone(),
	}
}
