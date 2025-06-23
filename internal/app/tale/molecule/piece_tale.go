package molecule

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/state"
	mycontrol "github.com/macabot/senet/internal/app/tale/control"
)

func TalePiece() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Piece",
		&state.State{
			Game: state.NewGame(),
		},
		func(s *state.State) *hypp.VNode {
			piece := s.Game.Board.FindPieceByID(1)
			props := molecule.PieceProps{
				Piece:         piece,
				Player:        s.Game.Turn,
				CanClick:      s.Game.CanClickOnPiece(s.Game.Turn, piece),
				DrawAttention: s.Game.PieceDrawsAttention(s.Game.Turn, piece.Position),
				Selected:      s.Game.PieceIsSelected(piece),
			}
			return molecule.Piece(props)
		},
	).WithControls(
		mycontrol.PlayerTurn(),
		control.NewCheckbox(
			"Can select",
			func(s *state.State, canSelect bool) hypp.Dispatchable {
				piece := s.Game.Board.FindPieceByID(1)
				canSelectPiece := s.Game.CanClickOnPiece(s.Game.Turn, piece)
				if canSelect == canSelectPiece {
					return s
				}

				if !canSelect {
					s.Game.Selected = nil
				}
				if s.Game.Turn == 0 {
					s.Game.TurnMode = state.IsPlayer0
				} else if s.Game.Turn == 1 {
					s.Game.TurnMode = state.IsPlayer1
				}
				s.Game.Sticks.HasThrown = canSelect
				return s
			},
			func(s *state.State) bool {
				return s.Game.CanClickOnPiece(s.Game.Turn, s.Game.Board.FindPieceByID(1))
			},
		),
		control.NewCheckbox(
			"Selected",
			func(s *state.State, selected bool) hypp.Dispatchable {
				if selected {
					s.Game.Selected = s.Game.Board.FindPieceByID(1)
				} else {
					s.Game.Selected = nil
				}
				return s
			},
			func(s *state.State) bool {
				return s.Game.PieceIsSelected(s.Game.Board.FindPieceByID(1))
			},
		),
		control.NewSelect(
			"Ability",
			func(s *state.State, ability state.PieceAbility) hypp.Dispatchable {
				piece := s.Game.Board.FindPieceByID(1)
				piece.Ability = ability
				return s
			},
			func(s *state.State) state.PieceAbility {
				return s.Game.Board.FindPieceByID(1).Ability
			},
			[]control.SelectOption[state.PieceAbility]{
				{Label: "Normal", Value: state.NormalPiece},
				{Label: "Protected", Value: state.ProtectedPiece},
				{Label: "Blocking", Value: state.BlockingPiece},
			},
		),
	)
}
