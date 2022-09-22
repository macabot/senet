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
		[4]int{0, 0, 0, 0},
		component.Sticks,
	).WithControls(
		fairy.NewButtonControl(
			"Throw",
			func(sticks [4]int) [4]int {
				return [4]int{
					flipStick(sticks[0]),
					flipStick(sticks[1]),
					flipStick(sticks[2]),
					flipStick(sticks[3]),
				}
			},
		),
	)
}
