package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func TopBar(s *state.State) *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": "top-bar",
		},
		rotateScreenButton(),
		Players(CreatePlayersProps(s)),
	)
}

func rotateScreenButton() *hypp.VNode {
	return html.Button(
		hypp.HProps{
			"class":   "rotate-screen",
			"onclick": dispatch.RotateScreenAction(),
		},
		ScreenRotationIcon(),
	)
}
