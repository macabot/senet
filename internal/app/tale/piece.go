package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func Piece() *fairytale.Tale[*state.State] {
	return fairytale.New(
		"Piece",
		&state.State{
			Game: state.NewGame(),
		},
		func(s *state.State) *hypp.VNode {
			piece := s.Game.Board.FindPieceByID(1)
			props := component.PieceProps{
				Piece:         piece,
				Player:        s.Game.Turn,
				CanClick:      s.Game.CanClickOnPiece(s.Game.Turn, piece),
				DrawAttention: s.Game.PiecesDrawAttention(s.Game.Turn),
				Selected:      s.Game.PieceIsSelected(piece),
			}
			return component.Piece(props)
		},
	).WithControls(
		control.NewSelect(
			"Player",
			func(s *state.State, player int) hypp.Dispatchable {
				s.Game.Turn = player
				return s
			},
			func(s *state.State) int {
				return s.Game.Turn
			},
			[]control.SelectOption[int]{
				{Label: "Player 1", Value: 0},
				{Label: "Player 2", Value: 1},
			},
		),
		control.NewCheckbox(
			"Can select",
			func(s *state.State, canSelect bool) *state.State {
				piece := s.Game.Board.FindPieceByID(1)
				canSelectPiece := s.Game.CanClickOnPiece(s.Game.Turn, piece)
				if canSelect == canSelectPiece {
					return s
				}

				if !canSelect {
					s.Game.Selected = nil
				}
				s.Game.HasTurn = canSelect
				s.Game.Sticks.HasThrown = canSelect
				return s
			},
			func(s *state.State) bool {
				return s.Game.CanClickOnPiece(s.Game.Turn, s.Game.Board.FindPieceByID(1))
			},
		),
		control.NewCheckbox(
			"Selected",
			func(s *state.State, selected bool) *state.State {
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
			func(s *state.State) int {
				return int(s.Game.Board.FindPieceByID(1).Ability)
			},
			[]control.SelectOption[state.PieceAbility]{
				{Label: "Normal", Value: state.NormalPiece},
				{Label: "Protected", Value: state.ProtectedPiece},
				{Label: "Blocking", Value: state.BlockingPiece},
			},
		),
	)
}
