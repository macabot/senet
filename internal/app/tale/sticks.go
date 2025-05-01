package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
	mycontrol "github.com/macabot/senet/internal/app/tale/control"
)

func Sticks() *fairytale.Tale[*state.State] {
	game := state.NewGame()
	game.TurnMode = state.IsPlayer0
	return fairytale.New(
		"Sticks",
		&state.State{Game: game},
		func(s *state.State) *hypp.VNode {
			gameCanThrow := s.Game.CanThrow()
			sticksCanThrow := s.Game.Sticks.CanThrow(s)
			return component.Sticks(component.SticksProps{
				Sticks:        s.Game.Sticks,
				DrawAttention: gameCanThrow && sticksCanThrow,
				NoValidMoves:  len(s.Game.ValidMoves) == 0,
				IsLoading:     gameCanThrow && !sticksCanThrow,
			})
		},
	).WithControls(
		mycontrol.Steps(),
		control.NewCheckbox(
			"No valid moves",
			func(s *state.State, noMoves bool) hypp.Dispatchable {
				if noMoves {
					s.Game.SetBoard(mycontrol.NoValidMovesBoard)
				} else {
					s.Game.SetBoard(state.NewBoard())
				}
				return s
			},
			func(s *state.State) bool {
				return len(s.Game.ValidMoves) == 0
			},
		),
		control.NewSelect(
			"Generator",
			func(s *state.State, kind state.SticksGeneratorKind) hypp.Dispatchable {
				s.Game.Sticks.GeneratorKind = kind
				return s
			},
			func(s *state.State) state.SticksGeneratorKind {
				return s.Game.Sticks.GeneratorKind
			},
			[]control.SelectOption[state.SticksGeneratorKind]{
				{Label: "Crypto", Value: state.CryptoSticksGeneratorKind},
				{Label: "Tutorial", Value: state.TutorialSticksGeneratorKind},
				{Label: "CommitmentScheme", Value: state.CommitmentSchemeGeneratorKind},
			},
		),
	)
}
