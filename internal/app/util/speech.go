package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/molecule"
	"github.com/macabot/senet/internal/app/state"
)

const delayStep = 50 * time.Millisecond

func speakWord(node *hypp.VNode, keyPrefix string, i int) *hypp.VNode {
	delay := time.Duration(i) * delayStep
	hProps := node.Props()
	hProps.Set("style", map[string]string{
		"animation-delay": fmt.Sprintf("%dms", delay.Milliseconds()),
	})
	hProps.Set("key", fmt.Sprintf("%s-%d", keyPrefix, i))
	return hypp.H(node.Tag(), hProps, node.Children()...)
}

func SpokenParagraph(text string, keyPrefix string) *hypp.VNode {
	var children []*hypp.VNode
	words := strings.SplitAfter(text, " ")
	for _, word := range words {
		for _, node := range ReplaceIcons(word) {
			children = append(children, speakWord(node, keyPrefix, len(children)))
		}
	}
	return html.P(
		hypp.HProps{"class": "spoken"},
		children...,
	)
}

func SpeechBubble(player int, bubble *state.SpeechBubble) *hypp.VNode {
	if bubble.DoNotRender {
		return nil
	}
	var speechVNodes []*hypp.VNode
	switch bubble.Kind {
	case state.DefaultSpeechBubble:
		speechVNodes = molecule.DefaultSpeechBubble()
	case state.TutorialStart:
		speechVNodes = molecule.TutorialStart(player)
	case state.TutorialGoal:
		speechVNodes = molecule.TutorialGoal(player)
	case state.TutorialPlayers1:
		speechVNodes = molecule.TutorialPlayers1(player)
	case state.TutorialPlayers2:
		speechVNodes = molecule.TutorialPlayers2()
	case state.TutorialBoard1:
		speechVNodes = molecule.TutorialBoard1(player)
	case state.TutorialBoard2:
		speechVNodes = molecule.TutorialBoard2(player)
	case state.TutorialBoard3:
		speechVNodes = molecule.TutorialBoard3(player)
	case state.TutorialSticks1:
		speechVNodes = molecule.TutorialSticks1(player)
	case state.TutorialSticks2:
		speechVNodes = molecule.TutorialSticks2(player)
	case state.TutorialSticks3:
		speechVNodes = molecule.TutorialSticks3()
	case state.TutorialMove:
		speechVNodes = molecule.TutorialMove()
	case state.TutorialMultiplemoves:
		speechVNodes = molecule.TutorialMultipleMoves(player)
	case state.TutorialTradingPlaces1:
		speechVNodes = molecule.TutorialTradingPlaces1(player)
	case state.TutorialTradingPlaces2:
		speechVNodes = molecule.TutorialTradingPlaces2(player)
	case state.TutorialTradingPlaces3:
		speechVNodes = molecule.TutorialTradingPlaces3(player)
	case state.TutorialTradingPlaces4:
		speechVNodes = molecule.TutorialTradingPlaces4()
	case state.TutorialBlockingPiece1:
		speechVNodes = molecule.TutorialBlockingPiece1(player)
	case state.TutorialBlockingPiece2:
		speechVNodes = molecule.TutorialBlockingPiece2()
	case state.TutorialReturnToStart1:
		speechVNodes = molecule.TutorialReturnToStart1(player)
	case state.TutorialReturnToStart2:
		speechVNodes = molecule.TutorialReturnToStart2(player)
	case state.TutorialReturnToStart3:
		speechVNodes = molecule.TutorialReturnToStart3()
	case state.TutorialMoveBackwards1:
		speechVNodes = molecule.TutorialMoveBackwards1(player)
	case state.TutorialMoveBackwards2:
		speechVNodes = molecule.TutorialMoveBackwards2()
	case state.TutorialNoMove1:
		speechVNodes = molecule.TutorialNoMove1(player)
	case state.TutorialNoMove2:
		speechVNodes = molecule.TutorialNoMove2()
	case state.TutorialOffTheBoard1:
		speechVNodes = molecule.TutorialOffTheBoard1(player)
	case state.TutorialOffTheBoard2:
		speechVNodes = molecule.TutorialOffTheBoard2()
	case state.TutorialOffTheBoard3:
		speechVNodes = molecule.TutorialOffTheBoard3()
	case state.TutorialEnd:
		speechVNodes = molecule.TutorialEnd()
	default:
		panic(fmt.Errorf("component not implemented for SpeechBubbleKind '%v'", bubble.Kind))
	}
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"speech-bubble": true,
				"closed":        bubble.Closed,
			},
		},
		speechVNodes...,
	)
}
