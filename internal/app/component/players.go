package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

type Player struct {
	Name   string
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
				Name:   s.Game.Players[0].Name,
				Points: s.Game.Board.Points(0),
			},
			{
				Name:   s.Game.Players[1].Name,
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
		playerTurnArrow(),
		player(1, props.Players[1], props.Turn == 1),
	)
}

func pointsIcon(points int) *hypp.VNode {
	switch points {
	case 0:
		return zeroPoints()
	case 1:
		return onePoint()
	case 2:
		return twoPoints()
	case 3:
		return threePoints()
	case 4:
		return fourPoints()
	case 5:
		return fivePoints()
	default:
		panic(fmt.Errorf("there exists no icon for %d points", points))
	}
}

func player(playerIndex int, player Player, hasTurn bool) *hypp.VNode {
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
	)
}
