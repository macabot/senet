package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func SignalingPage(s *state.State) *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - Player vs. Player")),
		html.Button(
			hypp.HProps{
				"class":   "signaling new-game",
				"onclick": dispatch.ToSignalingNewGamePageAction(),
			},
			hypp.Text("New game"),
		),
		html.Button(
			hypp.HProps{
				"class": "signaling join-game",
				// TODO "onclick"
			},
			hypp.Text("Join game"),
		),
		html.Button(
			hypp.HProps{
				"class":   "signaling back",
				"onclick": dispatch.ToStartPageAction(),
			},
			hypp.Text("Back"),
		),
	)
}

func SignalingNewGamePage(s *state.State) *hypp.VNode {
	offer := ""
	if s.Signaling != nil {
		offer = s.Signaling.Offer
	}
	return html.Main(
		hypp.HProps{
			"class": "signaling-page",
		},
		html.H1(nil, hypp.Text("Online - Player vs. Player")),
		html.P(nil, hypp.Text("Copy the text below and send it to your opponent.")),
		html.Textarea(nil, hypp.Text(offer)),
		html.Button(
			hypp.HProps{
				"class":   "signaling back",
				"onclick": dispatch.ToSignalingPageAction(),
			},
			hypp.Text("Back"),
		),
	)
}
