package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func StartPage(s *state.State) *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "start-page",
		},
		html.H1(nil, hypp.Text("Senet")),
		gameModes(),
	)
}

func gameModes() *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": "game-modes",
		},
		html.Button(
			hypp.HProps{
				"class":   "game-mode tutorial",
				"onclick": dispatch.ToTutorialAction(),
			},
			hypp.Text("Tutorial"),
		),
		html.Button(
			hypp.HProps{
				"class":   "game-mode local-pvp",
				"onclick": dispatch.ToLocalPlayerVsPlayerAction(),
			},
			hypp.Text("Local - Player vs. Player"),
		),
		html.Button(
			hypp.HProps{
				"class":    "game-mode local-pvb",
				"disabled": true,
			},
			hypp.Text("Local - Player vs. Bot"),
		),
		html.Button(
			hypp.HProps{
				"class":   "game-mode online-pvp",
				"onclick": dispatch.ToSignalingPageAction(),
			},
			hypp.Text("Online - Player vs. Player"),
		),
	)
}
