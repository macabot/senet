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
		Players(CreatePlayersProps(s)),
		html.Button(
			hypp.HProps{
				"class":      "menu-button",
				"onclick":    dispatch.ToggleMenuAction(),
				"aria-label": "menu button",
			},
			MenuIcon(),
		),
	)
}
