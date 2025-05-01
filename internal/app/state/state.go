package state

import (
	"encoding/json"
	"fmt"
)

type Page int

const (
	StartPage Page = iota
	SignalingPage
	WhoGoesFirstPage
	GamePage
	HomePage
	RulesPage
)

func (p Page) String() string {
	pages := [...]string{"Start", "Signaling", "WhoGoesFirst", "Game"}
	return pages[p]
}

func ToPage(s string) (Page, error) {
	var page Page
	switch s {
	case "Start":
		page = StartPage
	case "Signaling":
		page = SignalingPage
	case "WhoGoesFirst":
		page = WhoGoesFirstPage
	case "Game":
		page = GamePage
	default:
		return page, fmt.Errorf("invalid Page '%s'", s)
	}
	return page, nil
}

func (p Page) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *Page) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	var err error
	*p, err = ToPage(s)
	return err
}

type State struct {
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
