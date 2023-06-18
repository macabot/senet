package tale

import (
	"github.com/macabot/fairytale"
	"github.com/macabot/fairytale/control"
	"github.com/macabot/hypp"
	"github.com/macabot/senet/internal/app/component"
	"github.com/macabot/senet/internal/app/state"
	mycontrol "github.com/macabot/senet/internal/app/tale/control"
)

type BoardConfiguration struct {
	Label string
	Board *state.Board
}

type BoardConfigurations []BoardConfiguration

func (b BoardConfigurations) SelectOptions() []control.SelectOption[int] {
	options := make([]control.SelectOption[int], len(b))
	for i, configuration := range b {
		options[i] = control.SelectOption[int]{
			Label: configuration.Label,
			Value: i,
		}
	}
	return options
}

var boardConfigurations = BoardConfigurations{
	{
		Label: "New game",
		Board: state.NewBoard(),
	},
	{
		Label: "P1 - Protecting",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 9},
					&state.Piece{ID: 2, Position: 10},
					&state.Piece{ID: 3, Position: 5},
					&state.Piece{ID: 4, Position: 3},
					&state.Piece{ID: 5, Position: 1},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 8},
					&state.Piece{ID: 7, Position: 6},
					&state.Piece{ID: 8, Position: 4},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		},
	},
	{
		Label: "P2 - Protecting",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 9},
					&state.Piece{ID: 2, Position: 7},
					&state.Piece{ID: 3, Position: 5},
					&state.Piece{ID: 4, Position: 3},
					&state.Piece{ID: 5, Position: 1},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 8},
					&state.Piece{ID: 7, Position: 11},
					&state.Piece{ID: 8, Position: 4},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		},
	},
	{
		Label: "P1 - Blocking",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 9},
					&state.Piece{ID: 2, Position: 10},
					&state.Piece{ID: 3, Position: 11},
					&state.Piece{ID: 4, Position: 3},
					&state.Piece{ID: 5, Position: 1},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 8},
					&state.Piece{ID: 7, Position: 6},
					&state.Piece{ID: 8, Position: 4},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		},
	},
	{
		Label: "P2 - Blocking",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 9},
					&state.Piece{ID: 2, Position: 7},
					&state.Piece{ID: 3, Position: 5},
					&state.Piece{ID: 4, Position: 3},
					&state.Piece{ID: 5, Position: 1},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 8},
					&state.Piece{ID: 7, Position: 11},
					&state.Piece{ID: 8, Position: 12},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		},
	},
	{
		Label: "P1 - 2nd piece up",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 9},
					&state.Piece{ID: 2, Position: 12},
					&state.Piece{ID: 3, Position: 5},
					&state.Piece{ID: 4, Position: 3},
					&state.Piece{ID: 5, Position: 1},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 8},
					&state.Piece{ID: 7, Position: 6},
					&state.Piece{ID: 8, Position: 4},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		},
	},
	{
		Label: "P2 - 2nd piece up",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 9},
					&state.Piece{ID: 2, Position: 7},
					&state.Piece{ID: 3, Position: 5},
					&state.Piece{ID: 4, Position: 3},
					&state.Piece{ID: 5, Position: 1},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 8},
					&state.Piece{ID: 7, Position: 13},
					&state.Piece{ID: 8, Position: 4},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		},
	},
	{
		Label: "P1 - Protected by square",
		Board: &state.Board{
			PlayerPieces: [2]state.PiecesByPosition{
				state.NewPiecesByPosition(
					&state.Piece{ID: 1, Position: 9},
					&state.Piece{ID: 2, Position: 7},
					&state.Piece{ID: 3, Position: 25},
					&state.Piece{ID: 4, Position: 3},
					&state.Piece{ID: 5, Position: 1},
				),
				state.NewPiecesByPosition(
					&state.Piece{ID: 6, Position: 8},
					&state.Piece{ID: 7, Position: 6},
					&state.Piece{ID: 8, Position: 4},
					&state.Piece{ID: 9, Position: 2},
					&state.Piece{ID: 10, Position: 0},
				),
			},
		},
	},
}

func Board() *fairytale.Tale[*state.State] {
	game := state.NewGame()
	game.HasTurn = true
	return fairytale.New(
		"Board",
		&state.State{
			Game: game,
		},
		component.Board,
	).WithControls(
		control.NewSelect(
			"Configuration",
			func(s *state.State, option int) hypp.Dispatchable {
				if option == -1 {
					return s
				}
				board := boardConfigurations[option].Board
				s.Game.SetBoard(board)
				return s
			},
			func(s *state.State) int {
				for i, configuration := range boardConfigurations {
					if configuration.Board.Equal(s.Game.Board) {
						return i
					}
				}
				return -1
			},
			boardConfigurations.SelectOptions(),
		),
		control.NewCheckbox(
			"Has turn",
			func(s *state.State, hasTurn bool) *state.State {
				s.Game.HasTurn = hasTurn
				return s
			},
			func(s *state.State) bool {
				return s.Game.HasTurn
			},
		),
		mycontrol.PlayerTurn(),
		control.NewSelect(
			"Steps",
			func(s *state.State, steps int) hypp.Dispatchable {
				s.Game.SetSticks(state.SticksFromSteps(steps, steps != 0))
				return s
			},
			func(s *state.State) int {
				if !s.Game.Sticks.HasThrown {
					return 0
				}
				return s.Game.Sticks.Steps()
			},
			[]control.SelectOption[int]{
				{Label: "Not thrown", Value: 0},
				{Label: "1", Value: 1},
				{Label: "2", Value: 2},
				{Label: "3", Value: 3},
				{Label: "4", Value: 4},
				{Label: "6", Value: 6},
			},
		),
		control.NewSelect(
			"Selected",
			func(s *state.State, id int) hypp.Dispatchable {
				if id <= 0 {
					s.Game.SetSelected(nil)
				} else {
					s.Game.SetSelected(s.Game.Board.FindPieceByID(id))
				}
				return s
			},
			func(s *state.State) int {
				if s.Game.Selected == nil {
					return 0
				}
				return s.Game.Selected.ID
			},
			[]control.SelectOption[int]{
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
