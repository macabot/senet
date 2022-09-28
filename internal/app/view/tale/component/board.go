package component

import (
	"github.com/macabot/fairytale/fairy"
	"github.com/macabot/senet/internal/app/view/render/component"
	"github.com/macabot/senet/internal/app/view/state"
)

func BoardTale() *fairy.Tale {
	configuration := 0
	return fairy.NewTale(
		"Board",
		&state.State{
			Game: state.Game{
				Board: state.NewBoard(),
			},
		},
		component.Board,
	).WithControls(
		fairy.NewSelectControl(
			"Configuration",
			func(props *state.State, option int) *state.State {
				configuration = option
				props.Game.Board = state.NewBoard()
				switch option {
				case 0:
					// no-op
				case 1:
					props.Game.Board.Move(state.Position{2, 2}, 3)
				case 2:
					props.Game.Board.Move(state.Position{2, 3}, 5)
				case 3:
					props.Game.Board.Move(state.Position{2, 2}, 3)
					props.Game.Board.Move(state.Position{2, 4}, 6)
				case 4:
					props.Game.Board.Move(state.Position{2, 3}, 4)
					props.Game.Board.Move(state.Position{2, 5}, 7)
				}
				return props
			},
			func(_ *state.State) int {
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
			func(props *state.State, you int) *state.State {
				props.Game.You = you
				return props
			},
			func(props *state.State) int {
				return props.Game.You
			},
			[]fairy.SelectOption[int]{
				{Label: "Player 1", Value: 0},
				{Label: "Player 2", Value: 1},
			},
		),
		fairy.NewSelectControl(
			"Throw",
			func(props *state.State, throw int) *state.State {
				props.Game.Sticks = state.SticksFromThrow(throw, throw != 0)
				return props
			},
			func(props *state.State) int {
				if !props.Game.Sticks.HasThrown {
					return 0
				}
				return props.Game.Sticks.Value()
			},
			[]fairy.SelectOption[int]{
				{Label: "Not thrown", Value: 0},
				{Label: "1", Value: 1},
				{Label: "2", Value: 2},
				{Label: "3", Value: 3},
				{Label: "4", Value: 4},
				{Label: "6", Value: 6},
			},
		),
		fairy.NewSelectControl(
			"Selected",
			func(props *state.State, id int) *state.State {
				if id <= 0 {
					props.Game.Board.Selected = nil
				} else {
					pieces := props.Game.Board.PlayerPieces[0]
					if id >= 6 {
						pieces = props.Game.Board.PlayerPieces[1]
					}
					props.Game.Board.Selected = pieces.Find(state.ByID(id))
				}
				return props
			},
			func(props *state.State) int {
				if props.Game.Board.Selected == nil {
					return 0
				}
				return props.Game.Board.Selected.ID
			},
			[]fairy.SelectOption[int]{
				{Label: "Not Selected", Value: 0},
				{Label: "1", Value: 1},
				{Label: "2", Value: 2},
				{Label: "3", Value: 3},
				{Label: "4", Value: 4},
				{Label: "5", Value: 5},
				{Label: "6", Value: 6},
				{Label: "7", Value: 7},
				{Label: "8", Value: 8},
				{Label: "9", Value: 9},
				{Label: "10", Value: 10},
			},
		),
	)
}
