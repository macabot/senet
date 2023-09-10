package component

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func SpeechBubble(player int, bubble *state.SpeechBubble) *hypp.VNode {
	var speechVNodes []*hypp.VNode
	switch bubble.Kind {
	case state.DefaultSpeechBubble:
		speechVNodes = DefaultSpeechBubble()
	case state.TutorialStart:
		speechVNodes = TutorialStart(player)
	case state.TutorialPlayers1:
		speechVNodes = TutorialPlayers1(player)
	case state.TutorialPlayers2:
		speechVNodes = TutorialPlayers2()
	case state.TutorialGoal:
		speechVNodes = TutorialGoal(player)
	case state.TutorialBoard1:
		speechVNodes = TutorialBoard1(player)
	case state.TutorialBoard2:
		speechVNodes = TutorialBoard2(player)
	case state.TutorialBoard3:
		speechVNodes = TutorialBoard3(player)
	case state.TutorialSticks1:
		speechVNodes = TutorialSticks1(player)
	case state.TutorialSticks2:
		speechVNodes = TutorialSticks2(player)
	case state.TutorialSticks3:
		speechVNodes = TutorialSticks3()
	case state.TutorialMove:
		speechVNodes = TutorialMove()
	case state.TutorialMultiplemoves:
		speechVNodes = TutorialMultipleMoves(player)
	case state.TutorialTradingPlaces1:
		speechVNodes = TutorialTradingPlaces1(player)
	case state.TutorialTradingPlaces2:
		speechVNodes = TutorialTradingPlaces2(player)
	case state.TutorialTradingPlaces3:
		speechVNodes = TutorialTradingPlaces3(player)
	case state.TutorialTradingPlaces4:
		speechVNodes = TutorialTradingPlaces4()
	case state.TutorialBlockingPiece1:
		speechVNodes = TutorialBlockingPiece1(player)
	case state.TutorialBlockingPiece2:
		speechVNodes = TutorialBlockingPiece2()
	case state.TutorialReturnToStart1:
		speechVNodes = TutorialReturnToStart1(player)
	case state.TutorialReturnToStart2:
		speechVNodes = TutorialReturnToStart2(player)
	case state.TutorialReturnToStart3:
		speechVNodes = TutorialReturnToStart3()
	case state.TutorialMoveBackwards1:
		speechVNodes = TutorialMoveBackwards1(player)
	case state.TutorialMoveBackwards2:
		speechVNodes = TutorialMoveBackwards2()
	case state.TutorialNoMove1:
		speechVNodes = TutorialNoMove1(player)
	case state.TutorialNoMove2:
		speechVNodes = TutorialNoMove2()
	case state.TutorialEnd:
		speechVNodes = TutorialEnd()
	default:
		panic(fmt.Errorf("Component not implemented for SpeechBubbleKind %v", bubble.Kind))
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

func DefaultSpeechBubble() []*hypp.VNode {
	return []*hypp.VNode{
		html.P(nil, hypp.Textf("[Nothing to see here]")),
	}
}

func speechBubbleIcon(s string) *hypp.VNode {
	switch s {
	case "[blocking-icon]":
		return BlockingIcon()
	case "[no-move-icon]":
		return NoMoveIcon()
	case "[piece-0-icon]":
		return html.Span(hypp.HProps{"class": "piece-icon player-0"})
	case "[piece-1-icon]":
		return html.Span(hypp.HProps{"class": "piece-icon player-1"})
	case "[player-1-icon]":
		return html.Span(hypp.HProps{"class": "player-icon player-1"})
	case "[protected-icon]":
		return ProtectedIcon()
	case "[return-to-start-icon]":
		return ReturnToStartIcon()
	case "[square-invalid-icon]":
		return html.Span(hypp.HProps{"class": "square-icon invalid-destination"})
	case "[square-valid-icon]":
		return html.Span(hypp.HProps{"class": "square-icon valid-destination"})
	case "[sticks-1-icon]":
		return html.Span(
			hypp.HProps{"class": "sticks-icon"},
			html.Span(hypp.HProps{"class": "stick-icon white"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon black"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon black"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon black"}, nil),
		)
	case "[sticks-2-icon]":
		return html.Span(
			hypp.HProps{"class": "sticks-icon"},
			html.Span(hypp.HProps{"class": "stick-icon white"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon black"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon black"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon white"}, nil),
		)
	case "[sticks-3-icon]":
		return html.Span(
			hypp.HProps{"class": "sticks-icon"},
			html.Span(hypp.HProps{"class": "stick-icon white"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon black"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon white"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon white"}, nil),
		)
	case "[sticks-4-icon]":
		return html.Span(
			hypp.HProps{"class": "sticks-icon"},
			html.Span(hypp.HProps{"class": "stick-icon white"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon white"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon white"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon white"}, nil),
		)
	case "[sticks-6-icon]":
		return html.Span(
			hypp.HProps{"class": "sticks-icon"},
			html.Span(hypp.HProps{"class": "stick-icon black"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon black"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon black"}, nil),
			html.Span(hypp.HProps{"class": "stick-icon black"}, nil),
		)
	default:
		panic(fmt.Errorf("Speech bubble icon not implemented for '%s'.", s))
	}
}

var iconPattern = regexp.MustCompile(`\[[a-z0-9-]+-icon\]`)

func replaceIcons(text string) []*hypp.VNode {
	pairs := iconPattern.FindAllStringIndex(text, -1)
	var nodes []*hypp.VNode
	lastEnd := 0
	for _, pair := range pairs {
		start := pair[0]
		end := pair[1]

		if lastEnd < start {
			nodes = append(nodes, html.Span(nil, hypp.Text(text[lastEnd:start])))
		}
		nodes = append(nodes, speechBubbleIcon(text[start:end]))

		lastEnd = end
	}
	if lastEnd < len(text) {
		nodes = append(nodes, html.Span(nil, hypp.Text(text[lastEnd:])))
	}

	return nodes
}

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

func spokenParagraph(text string, keyPrefix string) *hypp.VNode {
	var children []*hypp.VNode
	words := strings.SplitAfter(text, " ")
	for _, word := range words {
		for _, node := range replaceIcons(word) {
			children = append(children, speakWord(node, keyPrefix, len(children)))
		}
	}
	return html.P(
		hypp.HProps{"class": "spoken"},
		children...,
	)
}
