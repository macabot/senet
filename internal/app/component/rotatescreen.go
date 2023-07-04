package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
)

func RotateScreen() *hypp.VNode {
	return html.Button(
		hypp.HProps{
			"class":   "rotate-screen",
			"onclick": dispatch.RotateScreenAction(),
		},
		ScreenRotationIcon(),
	)
}
