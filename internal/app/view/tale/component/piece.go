package component

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/senet/internal/app/view/render/component"
	"github.com/macabot/senet/internal/app/view/state"
)

func PieceTale() *fairytale.Tale {
	return fairytale.New(
		"Piece",
		component.PieceProps{
			Piece:     state.Piece{ID: 1, Position: 0},
			Player:    0,
			CanSelect: true,
			Moving:    false,
		},
		component.Piece,
	).WithControls(
		control.NewSelect(
			"Player",
			func(props component.PieceProps, player int) component.PieceProps {
				props.Player = player
				return props
			},
			func(props component.PieceProps) int {
				return props.Player
			},
			[]control.SelectOption[int]{
				{Label: "Player 1", Value: 0},
				{Label: "Player 2", Value: 1},
			},
		),
		control.NewCheckbox(
			"CanSelect",
			func(props component.PieceProps, canSelect bool) component.PieceProps {
				props.CanSelect = canSelect
				return props
			},
			func(props component.PieceProps) bool {
				return props.CanSelect
			},
		),
		control.NewCheckbox(
			"Moving",
			func(props component.PieceProps, moving bool) component.PieceProps {
				props.Moving = moving
				return props
			},
			func(props component.PieceProps) bool {
				return props.Moving
			},
		),
	)
}
