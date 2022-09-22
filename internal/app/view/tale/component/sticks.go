package component

import (
	"math/rand"

	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/view/render/component"
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

func SticksTale() *fairy.Tale {
	return fairy.NewTale(
		"Sticks",
		component.SticksProps{
			Sticks:   [4]int{0, 0, 0, 0},
			CanThrow: true,
		},
		component.Sticks,
	).WithControls(
		fairy.NewButtonControl(
			"Throw",
			func(props component.SticksProps) component.SticksProps {
				return component.SticksProps{
					Sticks: [4]int{
						flipStick(props.Sticks[0]),
						flipStick(props.Sticks[1]),
						flipStick(props.Sticks[2]),
						flipStick(props.Sticks[3]),
					},
					CanThrow: false,
				}
			},
		),
		fairy.NewCheckboxControl(
			"CanThrow",
			func(props component.SticksProps, canThrow bool) component.SticksProps {
				props.CanThrow = canThrow
				return props
			},
			func(props component.SticksProps) bool {
				return props.CanThrow
			},
		),
	)
}
