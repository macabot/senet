package state

// var SpeechBubbleFuncByName = map[string]func(int) *SpeechBubble{
// 	"TutorialStart":    TutorialStart,
// 	"TutorialPlayers1": TutorialPlayers1,
// 	"TutorialPlayers2": TutorialPlayers2,
// 	"TutorialGoal":     TutorialGoal,
// 	"TutorialBoard":    TutorialBoard,
// 	"TutorialEnd":      TutorialEnd,
// }

// func TutorialStart(player int) *SpeechBubble {
// 	return &SpeechBubble{
// 		Name: "TutorialStart",
// 		Elements: []*SpeechElement{
// 			{
// 				Kind:  TitleElement,
// 				Value: "Hello",
// 			},
// 			{
// 				Kind:  ParagraphElement,
// 				Value: "Hi, I'm the Tutor. Today I will teach you how to play Senet.",
// 			},
// 		},
// 		Button: SpeechButton{
// 			Text:    "Next",
// 			OnClick: NewSetSpeechBubbleAction(player, "TutorialPlayers1"),
// 		},
// 	}
// }

// func TutorialPlayers1(player int) *SpeechBubble {
// 	return &SpeechBubble{
// 		Name: "TutorialPlayers1",
// 		Elements: []*SpeechElement{
// 			{
// 				Kind:  TitleElement,
// 				Value: "Players - 1/2",
// 			},
// 			{
// 				Kind:  ParagraphElement,
// 				Value: "In the top you see the two players. You can hide or show the speech bubble of a player by clicking on that player.",
// 			},
// 		},
// 		Button: SpeechButton{
// 			Text:    "Next",
// 			OnClick: NewSetSpeechBubbleAction(player, "TutorialPlayers2"),
// 		},
// 	}
// }

// func TutorialPlayers2(player int) *SpeechBubble {
// 	return &SpeechBubble{
// 		Name: "TutorialPlayers2",
// 		Elements: []*SpeechElement{
// 			{
// 				Kind:  TitleElement,
// 				Value: "Players - 2/2",
// 			},
// 			replaceIcons(&SpeechElement{
// 				Kind:  ParagraphElement,
// 				Value: "Click on [player-icon] the Tutor to hide this speech bubble, then click on it again to show the speech bubble.",
// 			}),
// 		},
// 		Button: SpeechButton{
// 			Text:     "Next",
// 			Disabled: true,
// 		},
// 		EventListener: &EventListener{
// 			Event:  ClosePlayer2SpeechBubble,
// 			Action: *NewSetSpeechBubbleAction(player, "TutorialGoal"),
// 		},
// 	}
// }

// func TutorialGoal(player int) *SpeechBubble {
// 	return &SpeechBubble{
// 		Name: "TutorialGoal",
// 		Elements: []*SpeechElement{
// 			{
// 				Kind:  TitleElement,
// 				Value: "Goal",
// 			},
// 			{
// 				Kind:  ParagraphElement,
// 				Value: "The goal of Senet is to be the first player to move all of their pieces off the board.",
// 			},
// 		},
// 		Button: SpeechButton{
// 			Text:    "Next",
// 			OnClick: NewSetSpeechBubbleAction(player, "TutorialBoard"),
// 		},
// 	}
// }

// func TutorialBoard(player int) *SpeechBubble {
// 	return &SpeechBubble{
// 		Name: "TutorialBoard",
// 		Elements: []*SpeechElement{
// 			{
// 				Kind:  TitleElement,
// 				Value: "Board",
// 			},
// 			{
// 				Kind:  ParagraphElement,
// 				Value: "This is the board on which we play. Thie pieces flow in a Z shape from bottom right to top left.",
// 			},
// 		},
// 		Button: SpeechButton{
// 			Text: "Next",
// 			// TODO implement rest of tutorial bubbles.
// 		},
// 		OnCreate: &Action{
// 			Kind: ShowBoardFlow,
// 		},
// 	}
// }

// func TutorialEnd(player int) *SpeechBubble {
// 	return &SpeechBubble{
// 		Name: "TutorialEnd",
// 		Elements: []*SpeechElement{
// 			{
// 				Kind:  TitleElement,
// 				Value: "Good bye",
// 			},
// 			{
// 				Kind:  ParagraphElement,
// 				Value: "You now know how to play Senet. Go to the start page to start playing.",
// 			},
// 		},
// 		Button: SpeechButton{
// 			Text:    "Start page",
// 			OnClick: NewSetPage(StartPage),
// 		},
// 	}
// }
