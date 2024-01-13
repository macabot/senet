package state

import (
	"github.com/macabot/hypp"
)

type Page int

const (
	StartPage Page = iota
	SignalingPage
	WhoGoesFirstPage
	GamePage
)

type State struct {
	hypp.EmptyState
	PanicStackTrace    *string
	Game               *Game
	Page               Page
	ShowMenu           bool
	HideOrientationTip bool
	Signaling          *Signaling
	CommitmentScheme   CommitmentScheme
	TutorialIndex      int
}

func (s *State) Clone() *State {
	if s == nil {
		return nil
	}
	var panicStackTraceClone *string
	if s.PanicStackTrace != nil {
		p := *s.PanicStackTrace
		panicStackTraceClone = &p
	}
	return &State{
		PanicStackTrace:    panicStackTraceClone,
		Game:               s.Game.Clone(),
		Page:               s.Page,
		ShowMenu:           s.ShowMenu,
		HideOrientationTip: s.HideOrientationTip,
		Signaling:          s.Signaling.Clone(),
		CommitmentScheme:   s.CommitmentScheme.Clone(),
		TutorialIndex:      s.TutorialIndex,
	}
}
