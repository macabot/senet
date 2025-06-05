package template

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
)

func Head(title string) *hypp.VNode {
	return html.Head(
		nil,
		html.Meta(hypp.HProps{"charset": "utf-8"}),
		html.Meta(hypp.HProps{
			"name":    "viewport",
			"content": "width=device-width, initial-scale=1.0",
		}),
		html.Title(nil, hypp.Text(title)),
		html.Link(hypp.HProps{
			"rel":  "stylesheet",
			"href": "/senet.css",
		}),
	)
}
