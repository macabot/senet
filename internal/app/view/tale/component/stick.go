package component

import (
	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/view/render/component"
)

func StickTale() *fairy.Tale {
	return fairy.NewTale(
		"Stick",
		0,
		component.Stick,
	).WithControls(
		fairy.NewNumberInputControl(
			"Flips",
			func(_ int, flips int) int {
				return flips
			},
			func(flips int) int {
				return flips
			},
		),
	)
}
