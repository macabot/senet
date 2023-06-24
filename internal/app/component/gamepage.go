package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func GamePage(s *state.State) *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "game-page",
		},
		Players(CreatePlayersProps(s)),
		Board(s),
		Sticks(SticksProps{
			Sticks:        s.Game.Sticks,
			DrawAttention: s.Game.SticksDrawAttention(),
		}),
	)
}
