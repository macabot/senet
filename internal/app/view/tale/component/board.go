package component

import (
	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/model"
	"github.com/macabot/senet/internal/app/view/render/component"
	"github.com/macabot/senet/internal/app/view/state"
)

func BoardTale() *fairy.Tale {
	configuration := 0
	return fairy.NewTale(
		"Board",
		component.BoardProps{
			Board: model.NewBoard(),
			Meta:  state.Meta{},
		},
		component.Board,
	).WithControls(
		fairy.NewSelectControl(
			"Configuration",
			func(props component.BoardProps, option int) component.BoardProps {
				configuration = option
				props.Board = model.NewBoard()
				switch option {
				case 0:
					// no-op
				case 1:
					props.Board.Move(model.Position{2, 2}, 3)
				case 2:
					props.Board.Move(model.Position{2, 3}, 5)
				case 3:
					props.Board.Move(model.Position{2, 2}, 3)
					props.Board.Move(model.Position{2, 4}, 6)
				case 4:
					props.Board.Move(model.Position{2, 3}, 4)
					props.Board.Move(model.Position{2, 5}, 7)
				}
				return props
			},
			func(_ component.BoardProps) int {
				return configuration
			},
			[]fairy.SelectOption[int]{
				{Label: "New game", Value: 0},
				{Label: "P1- Protecting", Value: 1},
				{Label: "P2 - Protecting", Value: 2},
				{Label: "P1 - Blocking", Value: 3},
				{Label: "P2 - Blocking", Value: 4},
			},
		),
		fairy.NewSelectControl(
			"Player",
			func(props component.BoardProps, you int) component.BoardProps {
				props.Board.You = you
				return props
			},
			func(props component.BoardProps) int {
				return props.Board.You
			},
			[]fairy.SelectOption[int]{
				{Label: "Player 1", Value: 0},
				{Label: "Player 2", Value: 1},
			},
		),
	)
}
