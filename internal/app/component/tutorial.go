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
		html.P(nil, hypp.Text("Hi, I'm the Tutor. Today I will teach you how to play Senet.")),
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
		html.P(nil, hypp.Text("In the top you see the two players. You can hide or show the speech bubble of a player by clicking on that player.")),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialPlayers2),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialPlayers2(player int, buttonDisabled bool) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Players - 2/2")),
		html.P(nil, hypp.Text("Click on the Tutor [player-icon] to hide this speech bubble, then click on it again to show the speech bubble.")),
		html.Button(
			hypp.HProps{
				"onclick":  dispatch.SetSpeechBubbleKindAction(player, state.TutorialGoal),
				"disabled": buttonDisabled,
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialGoal(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Goal")),
		html.P(nil, hypp.Text("The goal of Senet is to be the first player to move all of their pieces off the board.")),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetSpeechBubbleKindAction(player, state.TutorialBoard),
			},
			hypp.Text("Next"),
		),
	}
}

func TutorialBoard(player int) []*hypp.VNode {
	return []*hypp.VNode{
		html.H3(nil, hypp.Text("Board")),
		html.P(nil, hypp.Text("This is the board on which we play. The pieces flow in a Z shape from bottom right to top left.")),
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
		html.P(nil, hypp.Text("You now know how to play Senet. Go to the start page to start playing.")),
		html.Button(
			hypp.HProps{
				"onclick": dispatch.SetPageAction(state.StartPage),
			},
			hypp.Text("Next"),
		),
	}
}
