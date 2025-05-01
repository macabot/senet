package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func GamePage(s *state.State) *hypp.VNode {
	gameCanThrow := s.Game.CanThrow()
	sticksCanThrow := s.Game.Sticks.CanThrow(s)
	return html.Main(
		hypp.HProps{
			"class": "game-page",
		},
		TopBar(s),
		Board(s),
		Sticks(SticksProps{
			Sticks:        s.Game.Sticks,
			DrawAttention: gameCanThrow && sticksCanThrow,
			NoValidMoves:  len(s.Game.ValidMoves) == 0,
			IsLoading:     gameCanThrow && !sticksCanThrow,
		}),
		Disconnected(s),
		GameOver(s),
		OrientationTip(s),
	)
}
