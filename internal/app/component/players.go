package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func Players(players [2]*state.Player, turn int) *hypp.VNode {
	return html.Section(
		hypp.HProps{
			"class": map[string]bool{
				"players":                        true,
				fmt.Sprintf("has-turn-%d", turn): true,
			},
		},
		player(0, players[0], turn == 0),
		playerTurnArrow(),
		player(1, players[1], turn == 1),
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

func player(playerIndex int, player *state.Player, hasTurn bool) *hypp.VNode {
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
