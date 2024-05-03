package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
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
			html.Button(
				hypp.HProps{
					"onclick": dispatch.GoToStartPage,
				},
				hypp.Text("Quit game"),
			),
			html.Button(
				hypp.HProps{
					"onclick": dispatch.ToggleMenu,
				},
				hypp.Text("Close"),
			),
		),
	)
}
