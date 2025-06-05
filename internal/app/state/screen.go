package state

import (
	"encoding/json"
	"fmt"
)

type Screen int

const (
	StartScreen Screen = iota
	OnlineScreen
	NewGameScreen
	JoinGameScreen
	WhoGoesFirstScreen
	GameScreen
)

func (s Screen) String() string {
	screens := [...]string{
		"StartScreen",
		"OnlineScreen",
		"NewGameScreen",
		"JoinGameScreen",
		"WhoGoesFirstScreen",
		"GameScreen",
	}
	return screens[s]
}

func ToScreen(s string) (Screen, error) {
	var screen Screen
	switch s {
	case "StartScreen":
		screen = StartScreen
	case "OnlineScreen":
		screen = OnlineScreen
	case "NewGameScreen":
		screen = NewGameScreen
	case "JoinGameScreen":
		screen = JoinGameScreen
	case "WhoGoesFirstScreen":
		screen = WhoGoesFirstScreen
	case "GameScreen":
		screen = GameScreen
	default:
		return screen, fmt.Errorf("invalid Screen '%s'", s)
	}
	return screen, nil
}

func (s Screen) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Screen) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	var err error
	*s, err = ToScreen(v)
	return err
}
