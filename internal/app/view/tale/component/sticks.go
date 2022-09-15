package component

import (
	"math/rand"

	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/view/render/component"
)

func SticksTale() *fairy.Tale {
	return fairy.NewTale(
		"Sticks",
		[4]bool{false, false, false, false},
		component.Sticks,
	).WithControls(
		fairy.NewButtonControl(
			"Throw",
			func(_ [4]bool) [4]bool {
				return [4]bool{
					rand.Float32() < 0.5,
					rand.Float32() < 0.5,
					rand.Float32() < 0.5,
					rand.Float32() < 0.5,
				}
			},
		),
	)
}
