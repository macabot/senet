package organism

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/dispatch"
)

func Menu() *hypp.VNode {
	return html.Aside(
		hypp.HProps{
			"class": "menu-wrapper",
		},
		html.Div(
			hypp.HProps{
				"class": "menu",
			},
			atom.Button("Quit game", dispatch.GoToStartScreen, nil),
			atom.Button("Close", dispatch.ToggleMenu, nil),
		),
	)
}
