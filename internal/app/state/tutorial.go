package state

var TutorialStart = &SpeechBubble{
	Name: "TutorialStart",
	Elements: []*SpeechElement{
		{
			Kind:  TitleElement,
			Value: "Hello",
		},
		{
			Kind:  ParagraphElement,
			Value: "Hi, I'm the Tutor. Today I will teach you how to play Senet.",
		},
	},
	Button: SpeechButton{
		Text: "Next",
		OnClick: &Action{
			Kind: GoToBubble,
			Data: TutorialPlayers1.Name,
		},
	},
}

var TutorialPlayers1 = &SpeechBubble{
	Name: "TutorialPlayers1",
	Elements: []*SpeechElement{
		{
			Kind:  TitleElement,
			Value: "Players - 1/2",
		},
		{
			Kind:  ParagraphElement,
			Value: "In the top you see the two players. You can hide or show the speech bubble of a player by clicking on that player.",
		},
	},
	Button: SpeechButton{
		Text: "Next",
		OnClick: &Action{
			Kind: GoToBubble,
			Data: TutorialPlayers2.Name,
		},
	},
}

var TutorialPlayers2 = &SpeechBubble{
	Name: "TutorialPlayers2",
	Elements: []*SpeechElement{
		{
			Kind:  TitleElement,
			Value: "Players - 2/2",
		},
		replaceIcons(&SpeechElement{
			Kind:  ParagraphElement,
			Value: "Click on [player-icon] the Tutor to hide this speech bubble, then click on it again to show the speech bubble.",
		}),
	},
	Button: SpeechButton{
		Text:     "Next",
		Disabled: true,
	},
	EventListener: &EventListener{
		Event: ClosePlayer2SpeechBubble,
		Action: Action{
			Kind: GoToBubble,
			Data: TutorialGoal.Name,
		},
	},
}

var TutorialGoal = &SpeechBubble{
	Name: "TutorialGoal",
	Elements: []*SpeechElement{
		{
			Kind:  TitleElement,
			Value: "Goal",
		},
		{
			Kind:  ParagraphElement,
			Value: "The goal of Senet is to be the first player to move all of their pieces off the board.",
		},
	},
	Button: SpeechButton{
		Text: "Next",
		OnClick: &Action{
			Kind: GoToBubble,
			Data: TutorialBoard.Name,
		},
	},
}

var TutorialBoard = &SpeechBubble{
	Name: "TutorialBoard",
	Elements: []*SpeechElement{
		{
			Kind:  TitleElement,
			Value: "Board",
		},
		{
			Kind:  ParagraphElement,
			Value: "This is the board on which we play. Thie pieces flow in a Z shape from bottom right to top left.",
		},
	},
	Button: SpeechButton{
		Text: "Next",
		OnClick: &Action{
			Kind: GoToBubble,
			Data: TutorialEnd.Name, // TODO implement rest of tutorial bubbles.
		},
	},
	OnCreate: Action{
		Kind: ShowBoardFlow,
	},
}

var TutorialEnd = &SpeechBubble{
	Name: "TutorialEnd",
	Elements: []*SpeechElement{
		{
			Kind:  TitleElement,
			Value: "Good bye",
		},
		{
			Kind:  ParagraphElement,
			Value: "You now know how to play Senet. Go to the start page to start playing.",
		},
	},
	Button: SpeechButton{
		Text: "Start page",
		OnClick: &Action{
			Kind: GoToPage,
			Data: StartPage,
		},
	},
}
