package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func GameOver(s *state.State) *hypp.VNode {
	if s.Game.Winner == nil {
		return nil
	}
	return html.Div(
		hypp.HProps{
			"class": "game-over",
		},
		html.P(
			nil,
			hypp.Textf("Winner: %s", s.Game.Players[*s.Game.Winner].Name),
		),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.GoToStartPage,
			},
			hypp.Text("Start page"),
		),
	)
}
