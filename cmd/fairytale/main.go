package main

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/app"
	"github.com/macabot/hypp"
	_ "github.com/macabot/hypp/jsd"
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
			Settings: fairytale.AdminSettings{
				IFrameSize:  fairytale.Size_iPhone_11_Pro,
				Orientation: fairytale.Landscape,
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
			tale.StartPage(),
			tale.WhoGoesFirstPage(),
		),
	)
}
