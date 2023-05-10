package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func Piece() *fairytale.Tale[*state.State] {
	props := component.PieceProps{
		Piece:         &state.Piece{ID: 1, Position: 9},
		Player:        0,
		CanSelect:     false,
		DrawAttention: false,
		Moving:        false,
		Selected:      false,
	}
	return fairytale.New(
		"Piece",
		nil,
		func(_ *state.State) *hypp.VNode {
			return component.Piece(props)
		},
	).WithControls(
		control.NewSelect(
			"Player",
			func(s *state.State, player int) hypp.Dispatchable {
				props.Player = player
				return s
			},
			func(_ *state.State) int {
				return props.Player
			},
			[]control.SelectOption[int]{
				{Label: "Player 1", Value: 0},
				{Label: "Player 2", Value: 1},
			},
		),
		control.NewCheckbox(
			"Can select",
			func(s *state.State, canSelect bool) *state.State {
				props.CanSelect = canSelect
				return s
			},
			func(_ *state.State) bool {
				return props.CanSelect
			},
		),
		control.NewCheckbox(
			"Draw attention",
			func(s *state.State, drawAttention bool) *state.State {
				props.DrawAttention = drawAttention
				return s
			},
			func(_ *state.State) bool {
				return props.DrawAttention
			},
		),
		control.NewCheckbox(
			"Moving",
			func(s *state.State, moving bool) *state.State {
				props.Moving = moving
				return s
			},
			func(_ *state.State) bool {
				return props.Moving
			},
		),
		control.NewCheckbox(
			"Selected",
			func(s *state.State, selected bool) *state.State {
				props.Selected = selected
				return s
			},
			func(_ *state.State) bool {
				return props.Selected
			},
		),
		control.NewSelect(
			"Ability",
			func(s *state.State, ability state.PieceAbility) hypp.Dispatchable {
				props.Piece.Ability = ability
				return s
			},
			func(_ *state.State) int {
				return int(props.Piece.Ability)
			},
			[]control.SelectOption[state.PieceAbility]{
				{Label: "Normal", Value: state.NormalPiece},
				{Label: "Protected", Value: state.ProtectedPiece},
				{Label: "Blocking", Value: state.BlockingPiece},
			},
		),
	)
}

/*
func Piece() *fairytale.Tale[*state.State] {
	s := &state.State{
		Game: state.NewGame(),
	}
	s.Game.SetHasTurn(true)
	return fairytale.New(
		"Piece",
		s,
		func(s *state.State) *hypp.VNode {
			player := s.Game.Turn
			props := component.PieceProps{
				Piece:     s.Game.Board.PlayerPieces[0][0],
				Player:    player,
				CanSelect: s.Game.CanSelect(player),
				Moving:    false, // TODO
			}
			return component.Piece(props)
		},
	).WithControls(
		control.NewSelect(
			"Player",
			func(s *state.State, player int) *state.State {
				s.Game.SetTurn(player)
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
		// control.NewCheckbox(
		// 	"CanSelect",
		// 	func(s *state.State, canSelect bool) component.PieceProps {
		// 		s.Game.CanSelect()
		// 		props.CanSelect = canSelect
		// 		return props
		// 	},
		// 	func(props component.PieceProps) bool {
		// 		return props.CanSelect
		// 	},
		// ),
		// control.NewCheckbox(
		// 	"Moving",
		// 	func(props component.PieceProps, moving bool) component.PieceProps {
		// 		props.Moving = moving
		// 		return props
		// 	},
		// 	func(props component.PieceProps) bool {
		// 		return props.Moving
		// 	},
		// ),
	)
}
*/
