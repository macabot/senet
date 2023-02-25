package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
)

func Piece() *fairytale.Tale[*state.State] {
	s := &state.State{
		Game: state.NewGame(),
	}
	s.Game.SetHasTurn(true)
	return fairytale.New(
		"Piece",
		s,
		func(s *state.State) *hypp.VNode {
			player := s.Game.Turn()
			props := component.PieceProps{
				Piece:     s.Game.Board().PlayerPieces[0][0],
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
				return s.Game.Turn()
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
