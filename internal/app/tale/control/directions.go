package control

import (
	"github.com/macabot/fairytale/control"
	"github.com/macabot/senet/internal/app/state"
)

func ShowDirections() *control.Checkbox[*state.State] {
	return control.NewCheckbox(
		"Show directions",
		func(s *state.State, showDirections bool) *state.State {
			s.Game.Board.ShowDirections = showDirections
			return s
		},
		func(s *state.State) bool {
			return s.Game.Board.ShowDirections
		},
	)
}
