package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
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
	)
}
