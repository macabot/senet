package component

import (
	"fmt"
	"strings"

	"github.com/macabot/hypp"
	htmlDriver "github.com/macabot/hypp/driver/html"
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
	if player.Speech != "" {
		bubble = speechBubble(player.Speech)
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

func speechBubble(speech string) *hypp.VNode {
	speechVNodes, err := htmlDriver.ParseFragment(strings.NewReader(speech), nil)
	if err != nil {
		speechVNodes = []*hypp.VNode{hypp.Text("[ERROR: Could not parse speech]")}
	}
	return html.P(
		hypp.HProps{
			"class": "speech-bubble",
		},
		speechVNodes...,
	)
}
