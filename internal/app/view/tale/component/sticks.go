package component

import (
	"math/rand"

	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/view/render/component"
	"github.com/macabot/senet/internal/app/view/state"
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

func SticksTale() *fairy.Tale[state.Sticks] {
	return fairy.NewTale(
		"Sticks",
		state.Sticks{
			Flips:     [4]int{0, 0, 0, 0},
			HasThrown: false,
		},
		component.Sticks,
	).WithControls(
		fairy.NewButtonControl(
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
		fairy.NewCheckboxControl(
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
