package component

import (
	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/model"
	"github.com/macabot/senet/internal/app/view/render/component"
)

func BoardTale() *fairy.Tale {
	configuration := 0
	return fairy.NewTale(
		"Board",
		model.NewBoard(),
		component.Board,
	).WithControls(
		fairy.NewSelectControl(
			"Configuration",
			func(board model.Board, option int) model.Board {
				configuration = option
				switch option {
				case 0:
					board = model.NewBoard()
				case 1:

				case 2:
				case 3:
				case 4:
				}
				return board
			},
			func(props model.Board) int {
				return configuration
			},
			[]fairy.SelectOption[int]{
				{Label: "New game", Value: 0},
				{Label: "P1- Protecting", Value: 1},
				{Label: "P2 - Protecting", Vaue: 2},
				{Label: "P1 - Blocking", Value: 3},
				{Label: "P2 - Blocking", Value: 4},
			}
		)
	)
}
