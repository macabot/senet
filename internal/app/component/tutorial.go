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
		html.H3(nil, hypp.Text("Players - 1/2")),
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
		html.H3(nil, hypp.Text("Players - 2/2")),
		spokenParagraph("Click on the Tutor [player-1-icon] to hide this speech bubble, then click on it again to show the speech bubble.", "TutorialPlayers2"),
	}
}

func TutorialGoal(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Goal")),
		spokenParagraph("The goal of Senet is to be the first player to move all of your pieces off the board.", "TutorialGoal"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialBoard1),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialBoard1(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Board - 1/3")),
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
		html.H3(nil, hypp.Text("Board - 2/3")),
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
		html.H3(nil, hypp.Text("Board - 3/3")),
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
		html.H3(nil, hypp.Text("Sticks - 1/3")),
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
		html.H3(nil, hypp.Text("Sticks - 2/3")),
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
		html.H3(nil, hypp.Text("Sticks - 3/3")),
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
		html.H3(nil, hypp.Text("Trading places - 1/3")),
		spokenParagraph("Let's change board to learn about trading the places of two pieces.", "TutorialTradingPlaces1"),
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
		html.H3(nil, hypp.Text("Trading places - 2/3")),
		spokenParagraph("A piece can move onto a square occupied by another player's piece, except if that piece is protected [protected-icon]. If it is not protected, then the pieces trade places. You are not allowed to trade places with a piece of the same color.", "TutorialTradingPlaces2"),
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
		html.H3(nil, hypp.Text("Trading places - 3/3")),
		spokenParagraph("A piece is protected [protected-icon] if at least one neighboring square - left, right, above or below - is occupied by a piece with the same color. A piece is also protected if it occupies a square that contains the protecting icon: [protected-icon].", "TutorialTradingPlaces3"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialEnd),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialEnd() []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Good bye")),
		spokenParagraph("You now know how to play Senet. Go to the start page to start playing.", "TutorialEnd"),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetPageAction(state.StartPage),
			},
			hypp.Text("Start page"),
		),
	}
}
