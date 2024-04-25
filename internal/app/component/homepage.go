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
			html.Span(nil, hypp.Text("Play ")),
			// TODO get size from environment variable.
			html.Span(hypp.HProps{"class": "wasm-size"}, hypp.Text("5 MB")),
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
