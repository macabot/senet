package molecule

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/dispatch"
)

func MenuButton() *hypp.VNode {
	return html.Button(
		hypp.HProps{
			"class":      "menu-button",
			"onclick":    dispatch.ToggleMenu,
			"aria-label": "menu button",
		},
		atom.MenuIcon(),
	)
}
