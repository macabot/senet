package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

type Player struct {
	*state.Player
	Points int
}

type PlayersProps struct {
	Players [2]Player
	Turn    int
}

func CreatePlayersProps(s *state.State) PlayersProps {
	return PlayersProps{
		Players: [2]Player{
			{
				Player: s.Game.Players[0],
				Points: s.Game.Board.Points(0),
			},
			{
				Player: s.Game.Players[1],
				Points: s.Game.Board.Points(1),
			},
		},
		Turn: s.Game.Turn,
	}
}

func Players(props PlayersProps) *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": map[string]bool{
				"players":                              true,
				fmt.Sprintf("has-turn-%d", props.Turn): true,
			},
		},
		player(0, props.Players[0], props.Turn == 0),
		PlayerTurnIcon(),
		player(1, props.Players[1], props.Turn == 1),
	)
}

func pointsIcon(points int) *hypp.VNode {
	switch points {
	case 0:
		return ZeroPointsIcon()
	case 1:
		return OnePointIcon()
	case 2:
		return TwoPointsIcon()
	case 3:
		return ThreePointsIcon()
	case 4:
		return FourPointsIcon()
	case 5:
		return FivePointsIcon()
	default:
		panic(fmt.Errorf("there exists no icon for %d points", points))
	}
}

func player(playerIndex int, player Player, hasTurn bool) *hypp.VNode {
	var bubble *hypp.VNode
	if player.SpeechBubble != nil {
		bubble = speechBubble(player.SpeechBubble)
	}
	return html.Div(
		hypp.HProps{
			"class": map[string]bool{
				"player":                              true,
				fmt.Sprintf("player-%d", playerIndex): true,
				"has-turn":                            hasTurn,
			},
		},
		html.Span(nil, hypp.Text(player.Name)),
		pointsIcon(player.Points),
		bubble,
	)
}

func speechBubble(bubble *state.SpeechBubble) *hypp.VNode {
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
	return html.Button(
		hypp.HProps{
			"disabled": button.Disabled,
			// TODO onclick
		},
		hypp.Text(button.Text),
	)
}
