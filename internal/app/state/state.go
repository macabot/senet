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
	PanicTrace         *string
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
	var panicTraceClone *string
	if s.PanicTrace != nil {
		p := *s.PanicTrace
		panicTraceClone = &p
	}
	return &State{
		PanicTrace:         panicTraceClone,
		Game:               s.Game.Clone(),
		Page:               s.Page,
		ShowMenu:           s.ShowMenu,
		HideOrientationTip: s.HideOrientationTip,
		Signaling:          s.Signaling.Clone(),
		CommitmentScheme:   s.CommitmentScheme.Clone(),
		TutorialIndex:      s.TutorialIndex,
	}
}
