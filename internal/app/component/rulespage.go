package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func RulesPage() *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "rules-page",
		},
		html.H1(nil, hypp.Text("Senet")),
		html.P(nil, replaceLinks("This page will explain the rules of Senet. You can also learn the rules with the [interactive-tutorial-link].")...),
		html.P(nil, hypp.Text("Senet is a two player game. The goal of Senet is to be the first player to move all of their pieces off the board.")),
		html.P(nil, replaceIcons("Below you see the board and pieces. Player 1 plays with the blue pieces [piece-0-icon]. Player 2 plays with the red pieces [piece-1-icon].")...),
		Board(&state.State{
			Game: state.NewGame(),
		}),
		html.P(nil, hypp.Text("The pieces move in a Z shape from bottom right to top left.")),
		Board(&state.State{
			Game: func() *state.Game {
				game := state.NewGame()
				game.Board.ShowDirections = true
				return game
			}(),
		}),
		html.P(nil, replaceIcons("Below the board we have the sticks. You can move a piece equal to the number of white sides. You can move a piece 1 step [sticks-1-icon], 2 steps [sticks-2-icon], 3 steps [sticks-3-icon] or 4 steps [sticks-4-icon]. If all sticks are showing the black side, you can move a piece 6 steps [sticks-6-icon].")...),
		Sticks(SticksProps{
			Sticks: &state.Sticks{
				Flips:     [4]int{1, 0, 0, 1},
				HasThrown: true,
			},
		}),
		html.P(nil, replaceIcons("To move a piece, click on it and then click on valid destination. A green square [square-valid-icon] is a valid destination. A red square [square-invalid-icon] is an invalid destination. The example below shows a blue piece [piece-0-icon] that has been selected and the corresponding valid destination [square-valid-icon].")...),
		// TODO pieces should not be clickable
		Board(&state.State{
			Game: func() *state.Game {
				game := state.NewGame()
				game.Selected = &state.Piece{Position: 9}
				game.Sticks = &state.Sticks{
					Flips:     [4]int{1, 0, 0, 1},
					HasThrown: true,
				}
				game.CalcValidMoves()
				return game
			}(),
		}),
	)
}
