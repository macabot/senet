package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func TutorialStart(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Hello")),
		spokenParagraph("Hi, I'm the Tutor. Today I will teach you how to play Senet.", "TutorialStart"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialPlayers1),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialPlayers1(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Players")),
		spokenParagraph("In the top you see the two players. You can hide or show the speech bubble of a player by clicking on that player.", "TutorialPlayers1"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialPlayers2),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialPlayers2() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Players")),
		spokenParagraph("Click on the [tutor-icon] to hide this speech bubble, then click on it again to show the speech bubble.", "TutorialPlayers2"),
	}
}

func TutorialBoard1(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Board")),
		spokenParagraph("Below the players is the board on which we play.", "TutorialBoard1"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialBoard2),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialBoard2(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Board")),
		spokenParagraph("At the bottom of the board are the pieces. You will play with the blue pieces [piece-0-icon]. I will play with the red pieces [piece-1-icon].", "TutorialBoard2"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialBoard3),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialBoard3(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Board")),
		spokenParagraph("The pieces move in a Z shape from bottom right to top left.", "TutorialBoard3"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialSticks1),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialSticks1(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Sticks")),
		spokenParagraph("At the bottom of the screen we find the sticks. You can move a piece equal to the number of white sides.", "TutorialSticks1"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialSticks2),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialSticks2(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Sticks")),
		spokenParagraph("You can move a piece 1 step [sticks-1-icon], 2 steps [sticks-2-icon], 3 steps [sticks-3-icon] or 4 steps [sticks-4-icon]. If all sticks are showing the black side, you can move a piece 6 steps [sticks-6-icon].", "TutorialSticks2"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialSticks3),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialSticks3() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Sticks")),
		spokenParagraph("Click on the sticks to throw the sticks.", "TutorialSticks3"),
	}
}

func TutorialMove() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Move")),
		spokenParagraph("You can now make your first move. Click on one of your pieces. A green square [square-valid-icon] is a valid destination. A red square [square-invalid-icon] is an invalid destination. Move a piece to a valid destination.", "TutorialMove"),
	}
}

func TutorialMultipleMoves(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Multiple moves")),
		spokenParagraph("If you throw 1 step [sticks-1-icon], 4 steps [sticks-4-icon] or 6 steps [sticks-6-icon], you may go again. This goes on until you throw 2 steps [sticks-2-icon] or 3 steps [sticks-3-icon]. Then your turn ends.", "TutorialMultipleMoves"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialTradingPlaces1),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialTradingPlaces1(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Trading places")),
		spokenParagraph("Let's change the board to learn about trading the places of two pieces.", "TutorialTradingPlaces1"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialTradingPlaces2),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialTradingPlaces2(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Trading places")),
		spokenParagraph("A piece can move to a square occupied by another player's piece, except if that piece is protected [protected-icon]. If not, the pieces trade places. You are not allowed to trade places with a piece of the same color.", "TutorialTradingPlaces2"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialTradingPlaces3),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialTradingPlaces3(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Trading places")),
		spokenParagraph("A piece is protected [protected-icon] if at least one neighboring square (left, right, above or below) is occupied by a piece with the same color or if it occupies a square with the protecting icon: [protected-icon].", "TutorialTradingPlaces3"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialTradingPlaces4),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialTradingPlaces4() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Trading places")),
		spokenParagraph("Move one of your pieces such that it trades places with one of my pieces.", "TutorialTradingPlaces4"),
	}
}

func TutorialBlockingPiece1(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Blocking piece")),
		spokenParagraph("Neighboring pieces that form a group of at least 3 pieces of the same color will block [blocking-icon] the movement of pieces of the other color. A piece that is blocking [blocking-icon] is also protected [protected-icon].", "TutorialBlockingPiece1"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialBlockingPiece2),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialBlockingPiece2() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Blocking piece")),
		spokenParagraph("Move one of your pieces. Note that you are not able to move over my blocking pieces [blocking-icon].", "TutorialBlockingPiece2"),
	}
}

func TutorialReturnToStart1(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Return to start")),
		spokenParagraph("The top row shows the return-to-start square [return-to-start-icon]. Let's change the board to learn about it.", "TutorialReturnToStart1"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialReturnToStart2),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialReturnToStart2(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Return to start")),
		spokenParagraph("If you move a piece to the return-to-start square [return-to-start-icon], then your piece is immediately moved to the start of the board: the first unoccupied square, starting in the bottom right of the board.", "TutorialReturnToStart2"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialReturnToStart3),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialReturnToStart3() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Return to start")),
		spokenParagraph("Move a piece to the return-to-start square [return-to-start-icon]. Close the speech bubble to see all available pieces.", "TutorialReturnToStart3"),
	}
}

func TutorialMoveBackwards1(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Move backwards")),
		spokenParagraph("If none of your pieces have a valid move forwards, then they must move backwards. When moving backwards, you are not allowed to trade places with another piece.", "TutorialMoveBackwards1"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialMoveBackwards2),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialMoveBackwards2() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Move backwards")),
		spokenParagraph("Move a piece backwards. Note that you are still not allowed to move a piece if it passes over another player's blocking piece [blocking-icon].", "TutorialMoveBackwards2"),
	}
}

func TutorialNoMove1(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("No move")),
		spokenParagraph("Sometimes no move is possible. Let's change the board to learn more.", "TutorialNoMove1"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialNoMove2),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialNoMove2() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("No move")),
		spokenParagraph("If none of your pieces have a valid move forwards and none of your pieces have a valid move backwards, then you must perform no move [no-move-icon]. Throw the sticks and perform no move [no-move-icon].", "TutorialNoMove2"),
	}
}

func TutorialOffTheBoard1(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Off the board")),
		spokenParagraph("A piece that is located in the top left square of the board will be moved off the board if all pieces of that color are located in the top row. Let's look at an example.", "TutorialOffTheBoard1"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialOffTheBoard2),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialOffTheBoard2() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Off the board")),
		spokenParagraph("All of your pieces are on the top row except for one. Move this piece to the top row. When all your pieces are in the top row, the piece on the top left square will be moved off the board.", "TutorialOffTheBoard2"),
	}
}

func TutorialOffTheBoard3() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Off the board")),
		spokenParagraph("The goal of Senet is to be the first player to move all of your pieces off the board. Keep playing until all of your pieces have been moved off the board.", "TutorialOffTheBoard3"),
	}
}

func TutorialEnd() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Great work!")),
		spokenParagraph("You now know how to play Senet. Go to the start page to start playing.", "TutorialEnd"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetPageAction(state.StartPage),
			},
			hypp.Text("Start page"),
		),
	}
}
