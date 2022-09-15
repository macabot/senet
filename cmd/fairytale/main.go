package main

import (
	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/view/tale/component"
)

func main() {
	fairy.Run(
		fairy.NewTree(
			fairy.NewBranch(
				"Components",
				component.BoardTale(),
				component.PieceTale(),
				component.StickTale(),
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
