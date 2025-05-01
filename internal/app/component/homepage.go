package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func HomePage() *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "home-page",
		},
		SenetHeader(),
		html.P(
			hypp.HProps{"class": "description"},
			hypp.Text("Senet is a two-player board game. Block your opponent's moves and be the first player to move all of their pieces off the board."),
		),
		homePageOptions(),
	)
}

func homePageOptions() *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": "links",
		},
		html.A(
			hypp.HProps{
				"href": "/play",
			},
			hypp.Text("Play"),
		),
		html.A(
			hypp.HProps{
				"href": "/rules",
			},
			hypp.Text("Rules"),
		),
		html.A(
			hypp.HProps{"href": "https://github.com/macabot/senet"},
			hypp.Text("Source code"),
		),
	)
}
