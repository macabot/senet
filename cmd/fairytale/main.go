package main

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/book"
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/view/tale/component"
	"github.com/macabot/senet/internal/app/view/tale/page"
)

func main() {
	book.Open(
		fairytale.NewTree(
			fairytale.NewBranch(
				"Components",
				component.BoardTale(),
				component.PieceTale(),
				component.StickTale(),
				component.SticksTale(),
			),
			fairytale.NewBranch(
				"Pages",
				page.GamePageTale(),
			),
		),
		[]*hypp.VNode{
			html.Link(hypp.HProps{
				"rel":  "stylesheet",
				"href": "http://localhost:8001/senet.css",
			}),
		},
	)
}
