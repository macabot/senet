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
		Piece:         state.Piece{ID: 1, Position: 9},
		Player:        0,
		CanSelect:     false,
		DrawAttention: false,
		Moving:        false,
		Selected:      false,
		PieceAbility:  state.NormalPiece,
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
			func(_ *state.State, player int) *state.State {
				props.Player = player
				return nil
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
			func(_ *state.State, canSelect bool) *state.State {
				props.CanSelect = canSelect
				return nil
			},
			func(_ *state.State) bool {
				return props.CanSelect
			},
		),
		control.NewCheckbox(
			"Draw attention",
			func(_ *state.State, drawAttention bool) *state.State {
				props.DrawAttention = drawAttention
				return nil
			},
			func(_ *state.State) bool {
				return props.DrawAttention
			},
		),
		control.NewCheckbox(
			"Moving",
			func(_ *state.State, moving bool) *state.State {
				props.Moving = moving
				return nil
			},
			func(_ *state.State) bool {
				return props.Moving
			},
		),
		control.NewCheckbox(
			"Selected",
			func(_ *state.State, selected bool) *state.State {
				props.Selected = selected
				return nil
			},
			func(_ *state.State) bool {
				return props.Selected
			},
		),
		control.NewSelect(
			"Ability",
			func(_ *state.State, ability state.PieceAbility) *state.State {
				props.PieceAbility = ability
				return nil
			},
			func(_ *state.State) int {
				return int(props.PieceAbility)
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
