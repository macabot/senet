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
		SenetHeader(),
		html.P(nil, html.A(hypp.HProps{"href": "/"}, hypp.Text("Home"))),
		html.P(nil, replaceLinks("This page will explain the rules of Senet. You can also learn the rules with the [interactive-tutorial-link].")...),
		html.Ul(
			nil,
			html.Li(nil, html.A(hypp.HProps{"href": "#goal"}, hypp.Text("Goal"))),
			html.Li(nil, html.A(hypp.HProps{"href": "#board"}, hypp.Text("Board"))),
			html.Li(nil, html.A(hypp.HProps{"href": "#sticks"}, hypp.Text("Sticks"))),
			html.Li(nil, html.A(hypp.HProps{"href": "#trading-places"}, hypp.Text("Trading places"))),
			html.Li(nil, html.A(hypp.HProps{"href": "#blocking-piece"}, hypp.Text("Blocking piece"))),
			html.Li(nil, html.A(hypp.HProps{"href": "#return-to-start"}, hypp.Text("Return to start"))),
			html.Li(nil, html.A(hypp.HProps{"href": "#move-backwards"}, hypp.Text("Move backwards"))),
			html.Li(nil, html.A(hypp.HProps{"href": "#no-move"}, hypp.Text("No move"))),
			html.Li(nil, html.A(hypp.HProps{"href": "#off-the-board"}, hypp.Text("Off the board"))),
		),

		html.H2(hypp.HProps{"id": "goal"}, hypp.Text("Goal")),
		html.P(nil, hypp.Text("Senet is a two player game. The goal of Senet is to be the first player to move all of their pieces off the board.")),

		html.H2(hypp.HProps{"id": "board"}, hypp.Text("Board")),
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
		html.P(nil, replaceIcons("To move a piece, click on it and then click on valid destination. A green square [square-valid-icon] is a valid destination. A red square [square-invalid-icon] is an invalid destination.")...),

		html.H2(hypp.HProps{"id": "sticks"}, hypp.Text("Sticks")),
		html.P(nil, replaceIcons("Below the board we have the sticks. You can move a piece equal to the number of white sides. You can move a piece 1 step [sticks-1-icon], 2 steps [sticks-2-icon], 3 steps [sticks-3-icon] or 4 steps [sticks-4-icon]. If all sticks are showing the black side, you can move a piece 6 steps [sticks-6-icon]. If you throw 1 step [sticks-1-icon], 4 steps [sticks-4-icon] or 6 steps [sticks-6-icon], you may go again. This goes on until you throw 2 steps [sticks-2-icon] or 3 steps [sticks-3-icon]. Then your turn ends.")...),
		Sticks(SticksProps{
			Sticks: &state.Sticks{
				Flips:     [4]int{1, 0, 0, 1},
				HasThrown: true,
			},
		}),

		html.H2(hypp.HProps{"id": "trading-places"}, hypp.Text("Trading places")),
		html.P(nil, replaceIcons("A piece can move to a square occupied by another player's piece, except if that piece is protected [protected-icon]. If not, the pieces trade places. You are not allowed to trade places with a piece of the same color. A piece is protected [protected-icon] if at least one neighboring square (left, right, above or below) is occupied by a piece with the same color or if it occupies a square with the protecting icon: [protected-icon].")...),

		html.H2(hypp.HProps{"id": "blocking-piece"}, hypp.Text("Blocking piece")),
		html.P(nil, replaceIcons("Neighboring pieces that form a group of at least 3 pieces of the same color will block [blocking-icon] the movement of pieces of the other color. A piece that is blocking [blocking-icon] is also protected [protected-icon].")...),

		html.H2(hypp.HProps{"id": "return-to-start"}, hypp.Text("Return to start")),
		html.P(nil, replaceIcons("The top row shows the return-to-start square [return-to-start-icon]. If you move a piece to the return-to-start square [return-to-start-icon], then your piece is immediately moved to the start of the board: the first unoccupied square, starting in the bottom right of the board.")...),

		html.H2(hypp.HProps{"id": "move-backwards"}, hypp.Text("Move backwards")),
		html.P(nil, replaceIcons("If none of your pieces have a valid move forwards, then they must move backwards. When moving backwards, you are not allowed to trade places with another piece. You are still not allowed to move a piece if it passes over another player's blocking piece [blocking-icon].")...),

		html.H2(hypp.HProps{"id": "no-move"}, hypp.Text("No move")),
		html.P(nil, replaceIcons("Sometimes no move is possible. If none of your pieces have a valid move forwards and none of your pieces have a valid move backwards, then you must perform no move [no-move-icon].")...),

		html.H2(hypp.HProps{"id": "off-the-board"}, hypp.Text("Off the board")),
		html.P(nil, replaceIcons("A piece that is located in the top left square of the board will be moved off the board if all pieces of that color are located in the top row. The goal of Senet is to be the first player to move all of their pieces off the board.")...),
	)
}
