package component

import (
	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/view/render/component"
)

func StickTale() *fairy.Tale {
	return fairy.NewTale(
		"Stick",
		false,
		component.Stick,
	).WithControls(
		fairy.NewCheckboxControl(
			"Up",
			func(_ bool, up bool) bool {
				return up
			},
			func(up bool) bool {
				return up
			},
		),
	)
}
