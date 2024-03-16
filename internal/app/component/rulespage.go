package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func RulesPage() *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "rules-page",
		},
		html.H1(nil, hypp.Text("Senet")),
		html.P(nil, hypp.Text("This page will explain the rules of Senet")),
		html.P(nil, hypp.Text("Senet is a two player game. The goal of Senet is to be the first player to move all of their pieces off the board.")),
		html.P(nil, hypp.Text("Below you see the board are pieces. Player 1 plays with the blue pieces [piece-0-icon].Player 2 plays with the red pieces [piece-1-icon].")),
	)
}
