package page

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/senet/internal/app/view/render/page"
	"github.com/macabot/senet/internal/app/view/state"
)

func GamePageTale() *fairytale.Tale {
	return fairytale.New(
		"GamePage",
		&state.State{
			Game: state.NewGame(),
		},
		page.GamePage,
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
