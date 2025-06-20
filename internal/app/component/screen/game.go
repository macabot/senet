package screen

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/organism"
	"github.com/macabot/senet/internal/app/state"
)

func Game(s *state.State) *hypp.VNode {
	gameCanThrow := s.Game.CanThrow()
	sticksCanThrow := s.Game.Sticks.CanThrow(s)
	var menu *hypp.VNode
	if s.ShowMenu {
		menu = organism.Menu()
	}
	return html.Main(
		hypp.HProps{
			"class": "game-page",
		},
		organism.TopBar(s),
		organism.Board(s),
		organism.Sticks(organism.SticksProps{
			Sticks:        s.Game.Sticks,
			DrawAttention: gameCanThrow && sticksCanThrow,
			NoValidMoves:  len(s.Game.ValidMoves) == 0,
			IsLoading:     gameCanThrow && !sticksCanThrow,
		}),
		organism.Disconnected(s),
		organism.GameOver(s),
		organism.OrientationTip(s),
		menu,
	)
}
