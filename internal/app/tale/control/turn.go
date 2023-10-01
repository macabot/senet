package control

import (
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/state"
)

func PlayerTurn() *control.Select[*state.State, int] {
	return control.NewSelect(
		"Player turn",
		func(s *state.State, player int) hypp.Dispatchable {
			s.Game.SetTurn(player)
			return s
		},
		func(s *state.State) int {
			if s.Game == nil {
				return -1
			}
			return s.Game.Turn
		},
		[]control.SelectOption[int]{
			{Label: "Player 1", Value: 0},
			{Label: "Player 2", Value: 1},
		},
	)
}
