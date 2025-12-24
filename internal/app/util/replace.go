package util

import (
	"fmt"
	"regexp"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/component/atom"
)

var (
	iconPattern = regexp.MustCompile(`\[[a-z0-9-]+-icon\]`)
	linkPattern = regexp.MustCompile(`\[[a-z0-9-]+-link]`)
)

func ReplaceIcons(text string) []*hypp.VNode {
	return replacePlaceholders(text, iconPattern, speechBubbleIcon)
}

func ReplaceLinks(text string) []*hypp.VNode {
	return replacePlaceholders(text, linkPattern, linkNode)
}

func replacePlaceholders(text string, pattern *regexp.Regexp, mapper func(string) *hypp.VNode) []*hypp.VNode {
	pairs := pattern.FindAllStringIndex(text, -1)
	var nodes []*hypp.VNode
	lastEnd := 0
	for _, pair := range pairs {
		start := pair[0]
		end := pair[1]

		if lastEnd < start {
			nodes = append(nodes, html.Span(nil, hypp.Text(text[lastEnd:start])))
		}
		nodes = append(nodes, mapper(text[start:end]))

		lastEnd = end
	}
	if lastEnd < len(text) {
		nodes = append(nodes, html.Span(nil, hypp.Text(text[lastEnd:])))
	}

	return nodes
}

func speechBubbleIcon(s string) *hypp.VNode {
	switch s {
	case "[blocking-icon]":
		return atom.BlockingIcon()
	case "[no-move-icon]":
		return atom.NoMoveIcon()
	case "[piece-0-icon]":
		return html.Span(hypp.HProps{"class": "piece-icon player-0"})
	case "[piece-1-icon]":
		return html.Span(hypp.HProps{"class": "piece-icon player-1"})
	case "[protected-icon]":
		return atom.ProtectedIcon()
	case "[return-to-start-icon]":
		return atom.ReturnToStartIcon()
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
	case "[tutor-icon]":
		return html.Span(
			hypp.HProps{"class": "player-icon player-1"},
			hypp.Text("Tutor"),
		)
	default:
		panic(fmt.Errorf("speech bubble icon not implemented for '%s'", s))
	}
}

func linkNode(s string) *hypp.VNode {
	switch s {
	case "[interactive-tutorial-link]":
		return atom.A("interactive tutorial", "/play", nil)
	default:
		panic(fmt.Errorf("link node not implemented for '%s'", s))
	}
}
