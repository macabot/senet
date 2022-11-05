package page

import (
	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/view/render/page"
	"github.com/macabot/senet/internal/app/view/state"
)

func GamePageTale() *fairy.Tale {
	return fairy.NewTale(
		"GamePage",
		&state.State{
			Game: state.NewGame(),
		},
		page.GamePage,
	).WithControls(
		fairy.NewCheckboxControl(
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
