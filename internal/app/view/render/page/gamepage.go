package page

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/view/render/component"
	"github.com/macabot/senet/internal/app/view/state"
)

func GamePage(props *state.State) *hypp.VNode {
	return html.Main(
		hypp.HProps{
			"class": "game-page",
		},
		component.Board(props),
		component.Sticks(props.Game.Sticks()),
	)
}
