package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/dispatch"
	"github.com/macabot/senet/internal/app/state"
)

func SpeechBubble(bubble *state.SpeechBubble) *hypp.VNode {
	speechVNodes := make([]*hypp.VNode, len(bubble.Elements))
	for i, element := range bubble.Elements {
		speechVNodes[i] = speechElement(element)
	}
	if bubble.Button.Text != "" {
		speechVNodes = append(speechVNodes, speechButton(bubble.Button))
	}
	return html.Div(
		hypp.HProps{
			"class": "speech-bubble",
		},
		speechVNodes...,
	)
}

var elementKindToVNodeFunc = map[state.SpeechElementKind]func(hypp.HProps, ...*hypp.VNode) *hypp.VNode{
	state.TitleElement:     html.H3,
	state.ParagraphElement: html.P,
	state.IconElement:      html.I,
}

func speechElement(element *state.SpeechElement) *hypp.VNode {
	if element.Kind == state.TextElement {
		return hypp.Text(element.Value)
	} else if element.Kind == state.IconElement {
		return html.I(nil, hypp.Text("�")) // TODO
	}
	vNodeFunc, ok := elementKindToVNodeFunc[element.Kind]
	if !ok {
		panic(fmt.Errorf("VNode func not implemented for SpeechElementKind %d", element.Kind))
	}
	childNodes := make([]*hypp.VNode, len(element.Children))
	for i, childElement := range element.Children {
		childNodes[i] = speechElement(childElement)
	}
	if element.Value != "" {
		childNodes = append(childNodes, hypp.Text(element.Value))
	}
	return vNodeFunc(nil, childNodes...)
}

func speechButton(button state.SpeechButton) *hypp.VNode {
	hProps := hypp.HProps{
		"disabled": button.Disabled,
	}
	if button.OnClick != nil {
		hProps["onclick"] = dispatch.SpeechBubbleAction(*button.OnClick)
	}
	return html.Button(
		hProps,
		hypp.Text(button.Text),
	)
}
