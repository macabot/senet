package page

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/component/template"
)

func Home() *hypp.VNode {
	return html.Html(
		hypp.HProps{"lang": "en"},
		template.Head("Senet"),
		html.Body(nil, homeMain()),
	)
}

func homeMain() *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "home-page",
		},
		molecule.SenetHeader(),
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
		atom.A("Play", "/play", nil),
		atom.A("Rules", "/rules", nil),
		atom.A("Source code", "https://github.com/macabot/senet", nil),
	)
}
