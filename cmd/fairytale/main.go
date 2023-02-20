package main

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/app"
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/tale"
)

func main() {
	app.Run(
		&app.Options{
			Assets: []*hypp.VNode{
				html.Link(hypp.HProps{
					"rel":  "stylesheet",
					"href": "http://localhost:8001/senet.css",
				}),
			},
		},
		fairytale.NewBundle(
			"Components",
			tale.BoardTale(),
			tale.PieceTale(),
			tale.StickTale(),
			tale.SticksTale(),
		),
		fairytale.NewBundle(
			"Pages",
			tale.GamePageTale(),
		),
	)
}
