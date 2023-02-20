package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/senet/internal/app/component"
)

func StickTale() *fairytale.Tale {
	return fairytale.New(
		"Stick",
		0,
		component.Stick,
	).WithControls(
		control.NewNumberInput(
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
