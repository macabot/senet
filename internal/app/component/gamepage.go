package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/state"
)

func GamePage(props *state.State) *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": map[string]bool{
				"game-page":    true,
				"focus-sticks": !props.Game.Sticks.HasThrown,
			},
		},
		Board(props),
		Sticks(SticksProps{
			Sticks:        props.Game.Sticks,
			DrawAttention: props.Game.SticksDrawAttention(),
		}),
	)
}
