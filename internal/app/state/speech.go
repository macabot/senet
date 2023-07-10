package state

import (
	"regexp"

	"github.com/macabot/senet/internal/pkg/clone"
)

type SpeechElementKind int

const (
	TextElement SpeechElementKind = iota
	TitleElement
	ParagraphElement
	IconElement
)

type SpeechElement struct {
	Kind     SpeechElementKind
	Value    string
	Children []*SpeechElement
}

type ActionKind int

const (
	SetSpeechBubble ActionKind = iota + 1
	SetPage
	ShowBoardFlow
)

type Action struct {
	Kind ActionKind
	Data any
}

func (a *Action) Clone() *Action {
	if a == nil {
		return nil
	}
	c := *a
	return &c
}

type SetSpeechBubbleData struct {
	Player           int
	SpeechBubbleName string
}

func NewSetSpeechBubbleAction(player int, speechBubbleName string) *Action {
	return &Action{
		Kind: SetSpeechBubble,
		Data: SetSpeechBubbleData{
			Player:           player,
			SpeechBubbleName: speechBubbleName,
		},
	}
}

func NewSetPage(page Page) *Action {
	return &Action{
		Kind: SetPage,
		Data: page,
	}
}

type Event int

const (
	ClosePlayer2SpeechBubble Event = iota + 1
)

type SpeechButton struct {
	Text     string
	Disabled bool
	OnClick  *Action
}

func (b SpeechButton) Clone() SpeechButton {
	return SpeechButton{
		Text:     b.Text,
		Disabled: b.Disabled,
		OnClick:  b.OnClick.Clone(),
	}
}

type EventListener struct {
	Event  Event
	Action Action
}

func (l *EventListener) Clone() *EventListener {
	if l == nil {
		return nil
	}
	return &EventListener{
		Event:  l.Event,
		Action: l.Action,
	}
}

type SpeechBubble struct {
	Name          string
	Elements      []*SpeechElement
	Button        SpeechButton
	EventListener *EventListener
	OnCreate      *Action
}

func (b *SpeechBubble) Clone() *SpeechBubble {
	if b == nil {
		return nil
	}
	return &SpeechBubble{
		Elements:      clone.Slice(b.Elements),
		Button:        b.Button.Clone(),
		EventListener: b.EventListener.Clone(),
		OnCreate:      b.OnCreate.Clone(),
	}
}

var iconPattern = regexp.MustCompile(`\[[a-z0-9-]+-icon\]`)

func replaceIcons(element *SpeechElement) *SpeechElement {
	pairs := iconPattern.FindAllStringIndex(element.Value, -1)
	var children []*SpeechElement
	lastEnd := 0
	for _, pair := range pairs {
		start := pair[0]
		end := pair[1]

		if lastEnd < start {
			children = append(children, &SpeechElement{
				Kind:  TextElement,
				Value: element.Value[lastEnd:start],
			})
		}
		children = append(children, &SpeechElement{
			Kind:  IconElement,
			Value: element.Value[start:end],
		})

		lastEnd = end
	}
	if lastEnd < len(element.Value) {
		children = append(children, &SpeechElement{
			Kind:  TextElement,
			Value: element.Value[lastEnd:],
		})
	}

	return &SpeechElement{
		Kind:     element.Kind,
		Children: children,
	}
}
