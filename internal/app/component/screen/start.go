package screen

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/dispatch"
)

func Start() *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "start-page",
		},
		molecule.SenetHeader(),
		gameModes(),
	)
}

func gameModes() *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": "game-modes",
		},
		atom.A("Home", "/", nil),
		atom.Button("Tutorial", dispatch.GoToTutorial, nil),
		atom.Button("Local - Player vs. Player", dispatch.GoToLocalPlayerVsPlayer, nil),
		atom.Button("Online - Player vs. Player", dispatch.GoToSignalingPage, nil),
		atom.A("Rules", "/rules", nil),
		atom.A("Source code", "https://github.com/macabot/senet", nil),
	)
}
