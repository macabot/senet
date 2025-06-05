package screen

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
)

func Online() *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - Player vs. Player")),
		html.Button(
			hypp.HProps{
				"class":   "cta",
				"onclick": dispatch.SetSignalingStepNewGameOffer,
			},
			hypp.Text("New game"),
		),
		html.Button(
			hypp.HProps{
				"class":   "cta",
				"onclick": dispatch.SetSignalingStepJoinGameOffer,
			},
			hypp.Text("Join game"),
		),
		html.Button(
			hypp.HProps{
				"class":   "signaling back",
				"onclick": dispatch.GoToStartPage,
			},
			hypp.Text("Back"),
		),
	)
}
