package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func Sticks() *fairytale.Tale {
	return fairytale.New(
		"Sticks",
		state.Sticks{
			Flips:     [4]int{0, 0, 0, 0},
			HasThrown: false,
		},
		component.Sticks,
	).WithControls(
		control.NewButton(
			"Throw",
			func(sticks state.Sticks) state.Sticks {
				return sticks.Throw()
			},
		),
		control.NewCheckbox(
			"Has thrown",
			func(props state.Sticks, hasThrown bool) state.Sticks {
				props.HasThrown = hasThrown
				return props
			},
			func(props state.Sticks) bool {
				return props.HasThrown
			},
		),
	)
}
