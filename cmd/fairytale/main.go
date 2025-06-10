package main

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/app"
	"github.com/macabot/hypp"
	_ "github.com/macabot/hypp/jsd"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/tale/molecule"
	"github.com/macabot/senet/internal/app/tale/organism"
	"github.com/macabot/senet/internal/app/tale/page"
	"github.com/macabot/senet/internal/app/tale/screen"
)

func main() {
	app.Run(
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
		fairytale.NewBundle(
			"Molecules",
			molecule.TalePiece(),
			molecule.TaleStick(),
		),
		fairytale.NewBundle(
			"Organisms",
			organism.TaleBoard(),
			organism.TalePlayers(),
			organism.TaleSticks(),
		),
		fairytale.NewBundle(
			"Screens",
			screen.TaleGame(),
			screen.TaleStart(),
			screen.TaleWhoGoesFirst(),
		),
		fairytale.NewBundle(
			"Pages",
			page.TaleHome(),
			page.TaleRules(),
		),
	)
}
