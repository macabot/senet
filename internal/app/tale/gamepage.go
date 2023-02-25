package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func GamePage() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"GamePage",
		&state.State{
			Game: state.NewGame(),
		},
		component.GamePage,
	).WithControls(
		control.NewCheckbox(
			"Has thrown",
			func(props *state.State, hasThrown bool) *state.State {
				props.Game.SetSticks(state.SticksFromSteps(6, hasThrown))
				return props
			},
			func(props *state.State) bool {
				return props.Game.Sticks().HasThrown
			},
		),
	)
}
