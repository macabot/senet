package state

type State struct {
	PanicStackTrace    *string
	Game               *Game
	Screen             Screen
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
		Screen:             s.Screen,
		ShowMenu:           s.ShowMenu,
		HideOrientationTip: s.HideOrientationTip,
		Signaling:          s.Signaling.Clone(),
		CommitmentScheme:   s.CommitmentScheme.Clone(),
		TutorialIndex:      s.TutorialIndex,
	}
}
