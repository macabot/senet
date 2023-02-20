package tale

import (
	"math/rand"

	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

const minFlips = 3

func flipStick(flips int) int {
	sign := 1
	if rand.Float32() < 0.5 {
		sign = -1
	}
	flips += minFlips * sign
	up := rand.Float32() < 0.5
	if (flips%2 == 0) == up {
		flips += sign
	}
	return flips
}

func SticksTale() *fairytale.Tale {
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
			func(props state.Sticks) state.Sticks {
				return state.Sticks{
					Flips: [4]int{
						flipStick(props.Flips[0]),
						flipStick(props.Flips[1]),
						flipStick(props.Flips[2]),
						flipStick(props.Flips[3]),
					},
					HasThrown: true,
				}
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
