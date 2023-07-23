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
	case state.TutorialStart:
		speechVNodes = TutorialStart(player)
	case state.TutorialPlayers1:
		speechVNodes = TutorialPlayers1(player)
	case state.TutorialPlayers2:
		speechVNodes = TutorialPlayers2(player, bubble.ButtonDisabled)
	case state.TutorialGoal:
		speechVNodes = TutorialGoal(player)
	case state.TutorialBoard:
		speechVNodes = TutorialBoard(player)
	case state.TutorialEnd:
		speechVNodes = TutorialEnd()
	default:
		panic(fmt.Errorf("Component not implemented for SpeechBubbleKind %v", bubble.Kind))
	}
	return html.Div(
		hypp.HProps{
			"class": "speech-bubble",
		},
		speechVNodes...,
	)
}

func speechBubbleIcon(s string) *hypp.VNode {
	switch s {
	case "[player-1-icon]":
		return html.Span(hypp.HProps{"class": "player-icon player-1"})
	default:
		// TODO replace s with proper icon
		return html.I(nil, hypp.Text("�"))
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
