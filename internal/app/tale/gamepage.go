package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func GamePage() *fairytale.Tale[*state.State] {
	game := state.NewGame()
	game.HasTurn = true
	return fairytale.New(
		"GamePage",
		&state.State{
			Game: game,
		},
		component.GamePage,
	).WithControls(
		control.NewCheckbox(
			"Has thrown",
			func(props *state.State, hasThrown bool) *state.State {
				sticks := props.Game.Sticks
				sticks.HasThrown = hasThrown
				props.Game.SetSticks(sticks)
				return props
			},
			func(props *state.State) bool {
				return props.Game.Sticks.HasThrown
			},
		),
		control.NewSelect(
			"Turn",
			func(s *state.State, turn int) hypp.Dispatchable {
				s.Game.SetTurn(turn)
				return s
			},
			func(s *state.State) int {
				return s.Game.Turn
			},
			[]control.SelectOption[int]{
				{Label: "Player 1", Value: 0},
				{Label: "Player 2", Value: 1},
			},
		),
	)
}
