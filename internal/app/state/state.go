package state

import "github.com/macabot/hypp"

type Page int

const (
	StartPage Page = iota
	SignalingPage
	SignalingNewGamePage
	GamePage
)

type State struct {
	hypp.EmptyState
	Game      *Game
	Page      Page
	Signaling *Signaling
}

func (s *State) Clone() *State {
	if s == nil {
		return nil
	}
	return &State{
		Game:      s.Game.Clone(),
		Page:      s.Page,
		Signaling: s.Signaling.Clone(),
	}
}
