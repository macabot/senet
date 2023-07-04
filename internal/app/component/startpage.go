package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func StartPage(s *state.State) *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "start-page",
		},
		RotateScreen(),
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
				"class": "game-mode tutorial",
			},
			hypp.Text("Tutorial"),
		),
		html.Button(
			hypp.HProps{
				"class":    "game-mode local-pvp",
				"disabled": true,
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
				"class":    "game-mode online-pvp",
				"disabled": true,
			},
			hypp.Text("Online - Player vs. Player"),
		),
	)
}
