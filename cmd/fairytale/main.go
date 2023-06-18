package main

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/app"
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
	"github.com/macabot/senet/internal/app/tale"
)

func main() {
	app.Run[*state.State](
		&app.Options{
			Assets: []*hypp.VNode{
				html.Link(hypp.HProps{
					"rel":  "stylesheet",
					"href": "/senet.css",
				}),
			},
		},
		fairytale.NewBundle[*state.State](
			"Components",
			tale.Board(),
			tale.Piece(),
			tale.Players(),
			tale.Stick(),
			tale.Sticks(),
		),
		fairytale.NewBundle[*state.State](
			"Pages",
			tale.GamePage(),
		),
	)
}
