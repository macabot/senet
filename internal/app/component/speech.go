package component

import (
	"fmt"
	"regexp"

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

var iconPattern = regexp.MustCompile(`\[[a-z0-9-]+-icon\]`)

// TODO use this function
func replaceIcons(text string) []*hypp.VNode {
	pairs := iconPattern.FindAllStringIndex(text, -1)
	var nodes []*hypp.VNode
	lastEnd := 0
	for _, pair := range pairs {
		start := pair[0]
		end := pair[1]

		if lastEnd < start {
			nodes = append(nodes, hypp.Text(text[lastEnd:start]))
		}
		nodes = append(nodes, html.I(nil, hypp.Text(text[start:end])))

		lastEnd = end
	}
	if lastEnd < len(text) {
		nodes = append(nodes, hypp.Text(text[lastEnd:]))
	}

	return nodes
}
